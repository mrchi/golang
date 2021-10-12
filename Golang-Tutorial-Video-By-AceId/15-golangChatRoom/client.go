package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	Name       string
	ServerIp   string
	ServerPort int
	conn       net.Conn
	flag       int
}

func (this *Client) Menu() bool {
	var flag int
	fmt.Println("1. 广播模式")
	fmt.Println("2. 私聊模式")
	fmt.Println("3. 修改用户名")
	fmt.Println("0. 退出")

	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		this.flag = flag
		return true
	} else {
		fmt.Println("输入不合法")
		return false
	}
}

// 更新用户名
func (this *Client) UpdateName() bool {
	fmt.Println(">>>请输入用户名：")
	fmt.Scanln(&this.Name)

	msg := fmt.Sprintf("rename|%v\n", this.Name)
	_, err := this.conn.Write([]byte(msg))
	if err != nil {
		fmt.Printf("conn.Write error, %v\n", err)
		return false
	} else {
		return true
	}
}

// 广播模式
func (this *Client) Broadcast() {
	var msg string
	for {
		fmt.Println(">>>请输入广播内容（exit 退出）：")
		fmt.Scanln(&msg)

		if msg == "exit" {
			break
		}

		if len(msg) != 0 {
			_, err := this.conn.Write([]byte(msg + "\n"))
			if err != nil {
				fmt.Printf("conn.Write error, %v\n", err)
				break
			}
		}
	}
}

// 查询用户
func (this *Client) SelectUser() {
	_, err := this.conn.Write([]byte("who\n"))
	if err != nil {
		fmt.Printf("conn.Write error, %v\n", err)
	}
}

// 私聊模式
func (this *Client) PrivateChat() {
	this.SelectUser()

	var remoteName string
	fmt.Println(">>>请输入聊天对象用户名（exit 退出）：")
	fmt.Scanln(&remoteName)

	if remoteName == "exit" {
		return
	}

	var msg string
	for {
		fmt.Println(">>>请输入私聊用户内容（exit 退出）：")
		fmt.Scanln(&msg)

		if msg == "exit" {
			break
		}

		if len(msg) != 0 {
			_, err := this.conn.Write([]byte(fmt.Sprintf("to|%v|%v\n", remoteName, msg)))
			if err != nil {
				fmt.Printf("conn.Write error, %v\n", err)
				break
			}
		}
	}
}

// 处理服务端发送的消息
func (this *Client) HandleResponse() {
	// 永久阻塞，将 this.conn 中的数据 copy 到 os.Stdout 中
	io.Copy(os.Stdout, this.conn)
}

func (this *Client) Run() {
	// 如果 flag 为 0 就退出
	for this.flag != 0 {
		// 如果输入不合法就一直循环
		for this.Menu() != true {
		}
		// 根据不同模式区分业务
		switch this.flag {
		case 1:
			fmt.Println("进入广播模式")
			this.Broadcast()
			break
		case 2:
			fmt.Println("进入私聊模式")
			this.PrivateChat()
			break
		case 3:
			fmt.Println("进入修改用户名模式")
			this.UpdateName()
			break
		}
	}
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
		flag:       999,
	}
	return &client
}

var serverIp string
var serverPort int

// 一般在 init 中定义命令行参数
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "服务器 IP")
	flag.IntVar(&serverPort, "port", 8000, "服务器端口号")
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)

	// 处理服务器消息
	go client.HandleResponse()

	client.Run()
}
