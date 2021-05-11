package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
	"sync"
)

// 分布式版 Raft

// 针对每个节点 id 和端口的封装类型
type nodeInfo struct {
	id   string
	port string
}

// Raft 声明节点对象类型 Raft
type Raft struct {
	node            nodeInfo
	mu              sync.Mutex // 锁
	me              int        // 节点编号
	currentTerm     int        // 当前任期
	votedFor        int        // 为那个节点投票
	state           int        // 3个状态,0:follower 1:candidate 2:leader
	timeout         int        // 超时时间
	lastMessageTime int64      // 发送最后一条数据的时间
	currentLeader   int        // 设置当前节点领导

	message     chan bool // 节点间发信息的通道
	electCh     chan bool // 选举的通道
	heartBeat   chan bool // 心跳信号的通道
	heartbeatRe chan bool // 子节点给主节点返回心跳信号的通道
}

// Leader 对象
type Leader struct {
	// 任期
	Term     int
	// leader 编号
	LeaderId int
}

// 设置节点个数
const raftCount = 2
var leader = Leader{0, -1}
// 存储缓存信息
var bufferMessage = make(map[string]string)
// 处理数据库信息
var mysqlMessage = make(map[string]string)
// 操作消息数组下标
var messageId = 1
// 用 nodeTable 存储每个节点中的键值对
var nodeTable map[string]string

func main() {
	// 终端接收来的是数组
	if len(os.Args) > 1 {
		// 接收终端输入的信号
		userId := os.Args[1]
		// 字符串转换整型
		id, _ := strconv.Atoi(userId)
		fmt.Println(id)
		// 定义节点 id 和端口号
		nodeTable = map[string]string {
			"1": ":8000",
			"2": ":8001",
		}
		// 封装 nodeInfo 对象
		node := nodeInfo{id: userId, port: nodeTable[userId]}
		// 创建节点对象
		rf := Make(id)
		// 确保每个新建立的节点都有端口对应
		// 127.0.0.1:8000
		rf.node = node
		// 注册 rpc
		go func() {
			// 注册 rpc，为了实现远程链接
			rf.raftRegisterRPC(node.port)
		}()
		if userId == "1" {
			go func() {
				// 回调方法
				http.HandleFunc("/req", rf.getRequest)
				fmt.Println("监听8080")
				if err := http.ListenAndServe("8080", nil); err != nil {
					fmt.Println(err)
					return
				}
			}()
		}
	}
	for {;
	}
}

var clientWriter http.ResponseWriter

func (r *Raft)getRequest(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()  // 解析请求参数
	if len(request.Form["age"]) > 0 {
		clientWriter = writer
		fmt.Println("主节点广播客户端请求age", request.Form["age"][0])

		param := Param{Msg: request.Form["age"][0], MsgId: strconv.Itoa(messageId)}
		messageId++
		if leader.LeaderId == r.me {
			r.sendMessageToOtherNodes(param)
		}else {
			// 将消息转发给 leader
			leaderId := nodeTable[strconv.Itoa(leader.LeaderId)]
			// 连接远程 rpc 服务
			rpc, err := rpc.DialHTTP("tcp", "127.0.0.1"+leaderId)
			if err != nil {
				log.Fatal("\nrpc 转发 server 错误：", leader.LeaderId, err)
			}
			var bo = false
			// 首先给 leader 传递
			err = rpc.Call("Raft.ForwardingMessage", param, &bo)
			if err != nil {
				log.Fatal("\nrpc 转发 server 错误：", leader.LeaderId, err)
			}
		}
	}
}

func (rf *Raft)sendMessageToOtherNodes(param Param) {
	bufferMessage[param.MsgId] = param.Msg
	// 只有领导者才能给其它服务器发送消息
	if rf.currentLeader == rf.me {
		var successCount = 0
		fmt.Println("领导者发送数据中....")
		go func() {
			rf.broadcast(param, "Raft.LogDataCopy", func(ok bool) {
				// 需要其他服务端回应
				rf.heartbeatRe <- ok
			})
		}()

		for i := 0; i < raftCount; i++ {
			fmt.Println("等待其他服务端回应")
			select {
			case ok := <-rf.message:
				if ok {
					successCount++
					if successCount >= raftCount/2 {
						rf.mu.Lock()
						rf.lastMessageTime=millisecond()
						mysqlMessage[param.MsgId] = bufferMessage[param.MsgId]
						delete(bufferMessage, param.MsgId)
						if clientWriter != nil {
							fmt.Fprintf(clientWriter, "OK")
						}
						fmt.Println("领导者发送数据结束")
					}
				}
			}
		}
	}
}

// 注册 Raft 对象，注册后的目的为确保每个节点 raft 可以远程接受
func (node *Raft)raftRegisterRPC(port string) {
	// 注册一个服务器
	rpc.Register(node)
	// 把服务绑定到 http 协议上
	rpc.HandleHTTP()
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("注册 RPC 服务失败", err)
	}
}

// Make 创建节点对象
func Make(me int) *Raft {
	rf := &Raft{}
	rf.me = me

}