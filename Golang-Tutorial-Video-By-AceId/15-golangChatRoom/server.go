package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
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
	// 创建 user
	user := NewUser(conn, this)

	// 用户上线
	user.Online()

	// 监听用户是否活跃的 channel
	isLive := make(chan bool)

	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			// 下线通知
			if n == 0 {
				user.Offline()
				return
			}
			// 错误处理
			if err != nil && err != io.EOF {
				fmt.Printf("Conn Read err: %v\n", err)
			}

			// 提取用户消息
			msg := string(buf[:n-1])
			// 用户处理消息
			user.DoMessage(msg)

			// 标记当前用户是活跃的
			isLive <- true
		}
	}()

	for {
		select {
		// 当前用户超时
		case <-time.After(time.Second * 5):
			user.ForceOffline()
			conn.Close()
			return
		// 当前用户是活跃的
		case <-isLive:
		}
	}
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
		fmt.Printf("连接 %v 建立成功\n", conn.RemoteAddr())

		// 业务处理
		go this.Handle(conn)
	}
}
