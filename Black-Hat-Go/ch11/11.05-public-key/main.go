package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
)

func main() {
	// 生成公钥和私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalln(err)
	}
	publicKey := &privateKey.PublicKey

	// 定义消息
	message := []byte("hello world")
	label := []byte("")

	// 加密
	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, label)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("cipherText = %x\n", cipherText)

	// 解密
	plaintText, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, cipherText, label)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("plaintText = %s\n", plaintText)

	// 计算签名
	h := sha256.New()
	h.Write(message)
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, h.Sum(nil), nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("signature = %x\n", signature)

	// 验证签名
	err = rsa.VerifyPSS(publicKey, crypto.SHA256, h.Sum(nil), signature, nil)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Signature verified")
	}
}
