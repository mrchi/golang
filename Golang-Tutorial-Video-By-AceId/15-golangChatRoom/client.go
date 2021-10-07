package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	Name       string
	ServerIp   string
	ServerPort int
	conn       net.Conn
}

// 创建客户端
func NewClient(serverIp string, serverPort int) *Client {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Printf("连接服务器失败: %v\n", err)
		return nil
	}

	fmt.Println("连接服务器成功")

	client := Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		conn:       conn,
	}
	return &client
}

var serverIp string
var serverPort int

// 一把在 init 中定义命令行参数
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "服务器 IP")
	flag.IntVar(&serverPort, "port", 8000, "服务器端口号")
}

func main() {
	flag.Parse()
	NewClient(serverIp, serverPort)

	select {}
}
