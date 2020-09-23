package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// 创建用户结构体类型
type Client struct {
	Name string
	Addr string
	C    chan string
}

// 创建全局 map 用来存储在线用户
var onlineMap map[string]Client

// 创建全局 channel 用来传递用户消息
var message = make(chan string)

// 用来组织需要发送给所有在线用户消息的消息内容结构，返回组织好是内容
func MakeMge(client Client, msg string) (buf string) {
	buf = "[" + client.Addr + "]" + client.Name + ":" + msg
	return
}

func WriteToClient(client Client, conn net.Conn) {
	// 监听用户自带 channel 上是否有消息，有则直接写给当前用户
	for msg := range client.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	// 创建一个 channel 判断用户是否为活跃状态
	hasData := make(chan bool)

	// 获取用户网络地址 ip+port, 并将其转换为字符串类型
	cltNetAddr := conn.RemoteAddr().String()
	// 创建新连接用户的结构体信息，默认 用户名是 Ip+port
	client := Client{cltNetAddr, cltNetAddr, make(chan string)}

	// 将新连接用户添加到在线用户 map 中, key: ip_port value: client 结构体
	onlineMap[cltNetAddr] = client

	// 创建专门用来给当前用户发送消息的 go 程
	go WriteToClient(client, conn)

	// 调用 MakeMsg 函数获取用户需要发送用户上线和消息的内容到全局 channel
	message <- MakeMge(client, "Login")

	// 创建一个 channel，用来判断用户退出状态
	isQuit := make(chan bool)

	// 创建一个匿名 go 程，用来处理用户发送的聊天信息内容
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Printf("检测到 %s 客户端已退出.......\n", client.Name)
				return
			}
			if err != nil {
				fmt.Println("conn.Read client msg error:", err)
				return
			}
			// 将收到的用户消息，保存在 msg 变量中 string 类型
			msg := string(buf[:n-1])
			// 判断用户输入的命令内容,提取在线用户列表发送给自己
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("onLine userList:\n"))
				// 遍历当前 map 获取在线的用户
				for _, user := range onlineMap {
					userInfo := user.Addr + ":" + user.Name + "\n"
					conn.Write([]byte(userInfo))
				}
				// =======================================================================
				// 判断用户发送了更改名称命令
			} else if len(msg) >= 8 && strings.Split(msg, "|")[0] == "rename" {
				newName := strings.Split(msg, "|")[1] // msg[8:]
				client.Name = newName                 // 修改结构体中的成员 Name
				onlineMap[cltNetAddr] = client        // 更新在线用户列表
				conn.Write([]byte("Change Name successful.......\n"))
				// =======================================================================
			} else {
				// 将读到的用户消息，写入到 message 管道中
				message <- MakeMge(client, msg)
			}
			hasData <- true
		}
	}()

	// 保证不退出
	for {
		// 监听 channel 上的数据流动
		select {
		case <-isQuit:
			close(client.C)
			// 将退出的用户从在线用户列表 onlineMap 中移除
			delete(onlineMap, cltNetAddr)
			message <- MakeMge(client, "LogOut") // 写入用户退出消息到全局 channel
			return
		case <-hasData:
			/*
				什么都不做，目的是为了重置下面 case 的计时器
				执行到这里则说明用户处于活跃状态，再次循环监
				听已达到重置计时器的目的
			*/
		case <-time.After(time.Second * 60):
			delete(onlineMap, cltNetAddr)
			message <- MakeMge(client, "Time Out Leaved")
			return
		}
	}
}

func Manager() {
	// 初始化 onlineMap
	onlineMap = make(map[string]Client)

	// 循环监听全局 channel 中是否有数据, 有数据存储至 msg, 无数据阻塞
	for {
		msg := <-message
		// 循环发送消息给所有在线用户, 要想执行，必须 msg := <-message 执行完解除阻塞
		for _, clnt := range onlineMap {
			clnt.C <- msg
		}
	}
}

func main() {
	// 创建监听 Socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}
	defer listener.Close()

	// 创建管理者 go 程，管理 map 和全局 channel
	go Manager()

	fmt.Println("服务器启动完成，等待客户端连接...................")
	// 循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error:", err)
			return
		}

		// 启动 go 程处理客户端数据请求
		go HandlerConnect(conn)
	}
}
