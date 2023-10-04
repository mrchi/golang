package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

const (
	MD5_HASH    = "5F4DCC3B5AA765D61D8327DEB882CF99"
	SHA256_HASH = "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"
)

func main() {
	file, err := os.Open("wordlist.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		md5Hash := fmt.Sprintf("%X", md5.Sum([]byte(word)))
		if md5Hash == MD5_HASH {
			log.Printf("[+] Password found (MD5): %s\n", word)
		}
		sha256Hash := fmt.Sprintf("%x", sha256.Sum256([]byte(word)))
		if sha256Hash == SHA256_HASH {
			log.Printf("[+] Password found (SHA256): %s\n", word)
		}
	}
}
