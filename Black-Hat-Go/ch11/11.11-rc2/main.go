package main

import (
	"crypto/cipher"
	"encoding/binary"
	"encoding/hex"
	"log"
	"os"
	"regexp"
	"sync"

	luhn "github.com/joeljunstrom/go-luhn"
	"github.com/mrchi/golang/Black-Hat-Go/ch11/11.11-rc2/rc2"
)

var numeric = regexp.MustCompile(`^\d{8}$`)

type CryptoData struct {
	block cipher.Block
	key   []byte
}

func generate(start, stop uint64, out chan<- *CryptoData, done <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := start; i <= stop; i++ {
			select {
			case <-done:
				return
			default:
				key := make([]byte, 8)
				binary.BigEndian.PutUint64(key, i)
				block, err := rc2.New(key[3:], 40)
				if err != nil {
					log.Fatalln(err)
				}
				data := &CryptoData{
					block: block,
					key:   key[3:],
				}
				out <- data
			}

		}
	}()
}

func decrypt(cipherText []byte, in <-chan *CryptoData, done chan struct{}, wg *sync.WaitGroup) {
	plainText := make([]byte, len(cipherText))
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range in {
			select {
			case <-done:
				return
			default:
				// 这里假设是采用的 ECB 模式
				// 同时假设密文为 16 字节，分开为两个 8 字节来处理
				data.block.Decrypt(plainText[:rc2.BlockSize], cipherText[:rc2.BlockSize])
				if numeric.Match(plainText[:rc2.BlockSize]) {
					data.block.Decrypt(plainText[rc2.BlockSize:], cipherText[rc2.BlockSize:])
					if numeric.Match(plainText[rc2.BlockSize:]) && luhn.Valid(string(plainText)) {
						log.Printf("Card [%s] found using key [%x]\n", plainText, data.key)
						close(done)
						return
					}
				}
			}
		}
	}()
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: xxx cipherText")
	}
	cipherText, err := hex.DecodeString(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	var producerWg, consumerWg sync.WaitGroup
	var min, max = uint64(0x0000000000), uint64(0xffffffffff)
	var producerCount, consumerCount = uint64(75), 5000
	var step = (max-min)/producerCount + 1

	done := make(chan struct{})
	work := make(chan *CryptoData, consumerCount)

	log.Println("Starting producers...")
	var start, stop = min, min + step
	for i := 0; i < int(producerCount); i++ {
		if stop > max {
			stop = max
		}
		generate(start, stop, work, done, &producerWg)
		stop += step
		start += step
	}
	log.Println("Producers started")

	log.Println("Starting consumers...")
	for i := 0; i < consumerCount; i++ {
		decrypt(cipherText, work, done, &consumerWg)
	}
	log.Println("Consumers started")

	producerWg.Wait()
	close(work)
	consumerWg.Wait()
	log.Println("Brute-force complete")
}
