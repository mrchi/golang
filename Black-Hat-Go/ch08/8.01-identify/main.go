package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}

	for _, device := range devices {
		fmt.Println(device.Name)
		for _, address := range device.Addresses {
			fmt.Printf("\tIP:      %s\n", address.IP)
			fmt.Printf("\tNetmask: %s\n", address.Netmask)
		}
	}
}
