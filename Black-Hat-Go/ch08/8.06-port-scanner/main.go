package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

const (
	TARGET             = "192.168.0.1"
	WORKER_COUNT       = 1000
	IFACE              = "en1"                                                   // Device
	SNAPLEN      int32 = 320                                                     // 快照长度（每个帧要捕获的数据量）
	PROMISC            = true                                                    // 是否以混杂模式运行
	TIMEOUT            = pcap.BlockForever                                       // 超时时间
	FILTER             = "tcp[13] == 0x11 or tcp[13] == 0x10 or tcp[13] == 0x18" // BPF 语法过滤器
)

func main() {
	ports := make(chan int, WORKER_COUNT)
	track := make(chan int)
	pcapResults := make(map[string]int)

	// 创建 goroutine worker
	for i := 0; i < WORKER_COUNT; i++ {
		go worker(ports, track)
	}

	// pcap 开始监听
	go capture(IFACE, TARGET, pcapResults)

	// 向 ports 中写入端口
	go func() {
		for i := 1; i <= 65535; i++ {
			ports <- i
		}
		close(ports)
	}()

	// 等待全部 worker 完成
	for i := 0; i < WORKER_COUNT; i++ {
		<-track
	}
	close(track)

	// 输出结果
	for port, confidence := range pcapResults {
		fmt.Printf("Port %s open: confidence %d\n", port, confidence)
	}
}

func worker(ports, track chan int) {
	for p := range ports {
		// debug 日志
		log.Printf("worker: scanning port %d\n", p)
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", TARGET, p), 1000*time.Millisecond)
		if err != nil {
			continue
		} else {
			time.Sleep(1 * time.Second)
			conn.Close()
		}
	}
	track <- 1
}

func capture(iface string, target string, results map[string]int) {
	handle, err := pcap.OpenLive(iface, SNAPLEN, PROMISC, TIMEOUT)
	if err != nil {
		log.Panicln(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(FILTER); err != nil {
		log.Panicln(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("Capturing packets")
	for packet := range source.Packets() {
		networkLayer := packet.NetworkLayer()
		if networkLayer == nil {
			continue
		}
		transportLayer := packet.TransportLayer()
		if transportLayer == nil {
			continue
		}

		srcHost := networkLayer.NetworkFlow().Src().String()
		srcPort := transportLayer.TransportFlow().Src().String()

		if srcHost != target {
			continue
		}

		results[srcPort] += 1
	}
}
