package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个 Server 对象
func NewServer(ip string, port int) *Server {
	server := Server{ip, port}
	return &server
}

// 处理业务
func (this *Server) Handle(conn net.Conn) {
	fmt.Println("连接建立成功")
}

// 启动服务器端口
func (this *Server) Start() {
	// 监听地址
	address := fmt.Sprintf("%s:%d", this.Ip, this.Port)
	listener, err := net.Listen("tcp4", address)
	if err != nil {
		fmt.Printf("net.Listen error: %v\n", err)
		return
	}
	fmt.Printf("服务器正在监听 %s...\n", address)

	// 关闭监听
	defer listener.Close()

	for {
		// 接受连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("listener.Accept error: %v\n", err)
			continue
		}

		// 业务处理
		go this.Handle(conn)
	}
}
