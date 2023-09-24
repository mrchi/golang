package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

const (
	IFACE         = "en1"             // Device
	SNAPLEN int32 = 1600              // 快照长度（每个帧要捕获的数据量）
	PROMISC       = false             // 是否以混杂模式运行
	TIMEOUT       = pcap.BlockForever // 超时时间
	FILTER        = "tcp and port 80" // BPF 语法过滤器
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}

	devFound := false
	for _, device := range devices {
		if device.Name == IFACE {
			devFound = true
		}
	}
	if !devFound {
		log.Panicln("Device not found")
	}

	handle, err := pcap.OpenLive(IFACE, SNAPLEN, PROMISC, TIMEOUT)
	if err != nil {
		log.Panicln(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(FILTER); err != nil {
		log.Panicln(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		// 取应用层包数据，这里也就是 HTTP 报文
		appLayer := packet.ApplicationLayer()
		if appLayer == nil {
			continue
		}
		// 取出 payload bytes，过滤是否含有关键字
		payload := appLayer.Payload()
		if bytes.Contains(payload, []byte("user")) || bytes.Contains(payload, []byte("pass")) {
			fmt.Println(string(payload))
		}
	}
}

// 测试命令：xh httpbin.org/get user==user pass==pass -v
