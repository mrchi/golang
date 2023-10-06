package main

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/mrchi/golang/Black-Hat-Go/ch13/imginject/pnglib"
	"github.com/mrchi/golang/Black-Hat-Go/ch13/imginject/xor"
)

const KEY = "gopher"

func main() {
	filePath := os.Args[2]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	inject, err := pnglib.NewPNGInject(file)
	if err != nil {
		log.Fatalln(err)
	}
	var c pnglib.PNGChunk
	offset := inject.ReadIENDChunk(&c)

	switch os.Args[1] {
	case "read":
		if c.Size == 0 {
			log.Println("Get nothing.")
		} else {
			originalData := xor.XorDecode(c.Data, KEY)
			log.Printf("Payload Original: %#2x\n", originalData)
			log.Printf("Payload Encode: %#2x\n", c.Data)
			log.Printf("Get: %s\n", originalData)
		}
	case "write":
		secret := []byte(os.Args[3])
		encodeSecret := xor.XorEncode(secret, KEY)
		buf := inject.WriteIENDChunk(offset, &encodeSecret)
		outputFilePath := strings.TrimSuffix(filePath, path.Ext(filePath)) + "-r" + path.Ext(filePath)
		os.WriteFile(outputFilePath, buf.Bytes(), 0644)
		log.Printf("Payload Original: %#2x\n", secret)
		log.Printf("Payload Encode: %#2x\n", encodeSecret)
		log.Printf("Write and saved to %s\n", outputFilePath)
	default:
		log.Fatalln("Usage: read xx.png or write xx.png secret")
	}
}
