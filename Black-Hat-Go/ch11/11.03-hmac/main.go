package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

var KEY = []byte("password")

func checkMAC(message, key, recvMac []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	calcMac := mac.Sum(nil)

	return hmac.Equal(calcMac, recvMac)
}

func main() {
	message := []byte("hello world")
	recvMac, _ := hex.DecodeString("8f5f355441dc2722900f292004f3d8a83245ff4d6e3078a5b77a4d7a921eeae9")

	if checkMAC(message, KEY, recvMac) {
		log.Println("EQUAL")
	} else {
		log.Println("NOT EQUAL")
	}
}
