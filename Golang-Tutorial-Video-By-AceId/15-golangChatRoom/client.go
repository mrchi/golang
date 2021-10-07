package main

import (
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

func main() {
	NewClient("127.0.0.1", 8000)

	select {}
}
