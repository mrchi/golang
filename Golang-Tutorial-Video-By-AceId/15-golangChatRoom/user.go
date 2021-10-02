package main

import "net"

type User struct {
	Name   string
	Addr   string
	Ch     chan string
	conn   net.Conn
	server *Server
}

// 创建用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := User{
		Name:   userAddr,
		Addr:   userAddr,
		Ch:     make(chan string),
		conn:   conn,
		server: server,
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

// 上线
func (this *User) Online() {
	// 用户加入到在线用户列表中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播上线消息
	this.server.BroadCast(this, "上线了")
}

// 下线
func (this *User) Offline() {
	// 用户从在线用户列表中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播下线消息
	this.server.BroadCast(this, "下线了")
}

// 处理正常消息
func (this *User) DoMessage(msg string) {
	this.server.BroadCast(this, msg)

}
