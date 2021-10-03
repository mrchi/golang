package main

import (
	"fmt"
	"net"
	"strings"
)

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
	if msg == "who" {
		// 查询在线用户
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := fmt.Sprintf("[%v]%v 在线...", user.Addr, user.Name)
			this.Ch <- onlineMsg
		}
		this.server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 重命名用户
		newName := strings.Split(msg, "|")[1]
		// 判断新名称是否已经被使用
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.Ch <- fmt.Sprintf("用户名 %v 已被占用", newName)
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.Ch <- fmt.Sprintf("已更新用户名：%v", newName)
		}
	} else {
		this.server.BroadCast(this, msg)
	}
}
