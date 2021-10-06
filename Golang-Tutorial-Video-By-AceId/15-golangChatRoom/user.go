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

// 发送消息给客户端
func (this *User) SendMessage(message string) {
	this.conn.Write([]byte(message + "\n"))
}

// 监听当前 User channel 的方法，如果有消息就发给客户端
func (this *User) ListenMessage() {
	for {
		this.SendMessage(<-this.Ch)
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

// 被强制下线
func (this *User) ForceOffline() {
	this.SendMessage("你已超时，强制下线")
	close(this.Ch)
}

// 处理正常消息
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询在线用户
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := fmt.Sprintf("[%v]%v 在线...", user.Addr, user.Name)
			this.SendMessage(onlineMsg)
		}
		this.server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 重命名用户
		newName := strings.Split(msg, "|")[1]
		// 判断新名称是否已经被使用
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMessage(fmt.Sprintf("用户名 %v 已被占用", newName))
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMessage(fmt.Sprintf("已更新用户名：%v", newName))
		}
	} else if len(msg) > 3 && msg[:3] == "to|" {
		// 发送私聊消息，消息格式 to|zhang3|hello

		// 获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMessage("消息格式不正确")
			return
		}
		// 根据用户名获取到 User 对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMessage("该用户名不存在")
			return
		}
		// 获取消息内容并发送
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMessage("消息内容为空")
			return
		}
		remoteUser.SendMessage(fmt.Sprintf("[From %v]%v", this.Name, content))
	} else {
		this.server.BroadCast(this, msg)
	}
}
