package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"log"
)

func pad(buf []byte) []byte {
	// pcks#7
	bufLength := len(buf)
	paddingLength := aes.BlockSize - bufLength%aes.BlockSize
	result := make([]byte, bufLength+paddingLength)
	copy(result, buf)
	copy(result[bufLength:], bytes.Repeat([]byte{byte(paddingLength)}, paddingLength))
	return result
}

func unpad(buf []byte) []byte {
	// pcks#7
	paddingCount := int(buf[len(buf)-1])
	return buf[:len(buf)-paddingCount]
}

func encrypt(plainText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	plainText = pad(plainText)
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	copy(cipherText, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return cipherText, nil
}

func decrypt(cipherText, key []byte) ([]byte, error) {
	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("invalid cipher text length: too short")
	}
	if len(cipherText)%aes.BlockSize != 0 {
		return nil, errors.New("invalid cipher text length: not a multiple of blocksize")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	mode.CryptBlocks(plainText, cipherText)

	return unpad(plainText), nil
}

func main() {
	cipherText := []byte{
		0x69, 0x76, 0x69, 0x76, 0x69, 0x76, 0x69, 0x76,
		0x69, 0x76, 0x69, 0x76, 0x69, 0x76, 0x69, 0x76,
		0x22, 0x13, 0x5a, 0x3c, 0x61, 0x79, 0x20, 0x63,
		0x1e, 0x33, 0x2a, 0x1d, 0xd0, 0xb3, 0x31, 0x3c,
	}
	key := "passwordpassword"
	log.Printf("cipherText = %x\n", cipherText)

	plainText, err := decrypt(cipherText, []byte(key))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("plainText = %s\n", plainText)

	anotherCipherText, err := encrypt(plainText, []byte(key))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("anotherCipherText = %x\n", anotherCipherText)
}
