package main

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

const BCRYPT_HASH = "$2a$10$o1.L42pP2FPxmDQK305RUuabvX4c8uushC6pvMwold7DUOO/E2CWe"

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: bcrypt password")
	}
	password := os.Args[1]

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Hash = %s\n", hash)

	err = bcrypt.CompareHashAndPassword([]byte(BCRYPT_HASH), []byte(password))
	if err != nil {
		log.Println("[!] Authentication failed")
	} else {
		log.Println("[+] Authentication successful")
	}
}
