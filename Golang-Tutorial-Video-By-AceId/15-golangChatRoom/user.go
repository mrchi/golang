package main

import "net"

type User struct {
	Name string
	Addr string
	Ch   chan string
	conn net.Conn
}

// 创建用户
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := User{
		Name: userAddr,
		Addr: userAddr,
		Ch:   make(chan string),
		conn: conn,
	}

	// 启动监听 goroutine
	go user.ListenMessage()

	return &user
}

// 监听当前 User channel 的方法，如果有消息就发给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.Ch
		this.conn.Write([]byte(msg + "\n"))
	}
}
