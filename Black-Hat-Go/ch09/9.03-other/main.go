package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	// Python hex(x)
	fmt.Printf("%#x\n", 255)

	// Python ord(c)
	fmt.Printf("%#v\n", rune('d'))

	// Python chr(i)
	fmt.Printf("%#v\n", string(rune(100)))

	// Python struct.pack()
	data := []any{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	buf := new(bytes.Buffer)
	for _, v := range data {
		if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	result := buf.Bytes()
	fmt.Printf("%#v\n", result)

	// Python struct.unpack()
	var rawData struct {
		A uint16
		B int8
		C uint8
	}
	r := bytes.NewReader(result)
	if err := binary.Read(r, binary.LittleEndian, &rawData); err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("%d %d %d\n", rawData.A, rawData.B, rawData.C)

	// some formats -> byte slice
	fmt.Printf("%#v\n", []byte("\xfc\xe8\x82"))

	result, err := hex.DecodeString("fce882")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("%#v\n", result)

	fmt.Printf("%#v\n", []byte{0xfc, 0xe8, 0x82})

	result, err = base64.StdEncoding.DecodeString("/OiC")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("%#v\n", result)
}
