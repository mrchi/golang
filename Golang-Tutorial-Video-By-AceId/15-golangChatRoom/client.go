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
			break
		case 2:
			fmt.Println("进入私聊模式")
			break
		case 3:
			fmt.Println("进入修改用户名模式")
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

	client.Run()
}
