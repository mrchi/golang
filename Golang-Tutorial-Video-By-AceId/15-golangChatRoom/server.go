package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的 channel
	Message chan string
}

// 监听 Server.Message 广播消息 channel 的 goroutine，把消息发送给全部的在线 user
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		this.mapLock.Lock()
		for _, user := range this.OnlineMap {
			user.Ch <- msg
		}
		this.mapLock.Unlock()
	}
}

// 创建一个 Server 对象
func NewServer(ip string, port int) *Server {
	server := Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return &server
}

// 广播消息
func (this *Server) BroadCast(user *User, msg string) {
	message := fmt.Sprintf("[%v]%v:%v", user.Addr, user.Name, msg)
	this.Message <- message
}

// 处理业务
func (this *Server) Handle(conn net.Conn) {
	fmt.Println("连接建立成功")

	// 创建 user
	user := NewUser(conn)

	// 用户加入到在线用户列表中
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	// 广播消息
	this.BroadCast(user, "上线了")

	select {}
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

	// 启动监听 Message 的 goroutine
	go this.ListenMessage()

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
