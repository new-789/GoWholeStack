package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

// Raft 选举实现 Demo

// 1. 实现先节点选举
// 2. 改造代码成分布式选举代码，加入 RPC 调用
// 3. 演示完整代码，自动选主、日志复制

// 定义3个节点，常量
const raftCount = 3

// Leader 声明 leader 对象
type Leader struct {
	Term     int // 任期
	LeaderId int // Leader 编号

}

// Raft 声明 raft
type Raft struct {
	mu              sync.Mutex // 锁
	me              int        // 节点编号
	currentTerm     int        // 当前任期
	votedFor        int        // 为那个节点秃瓢
	state           int        // 3个状态,0:follower 1:candidate 2:leader
	lastMessageTime int64      // 发送最后一条数据的时间
	currentLeader   int        // 设置当前节点领导
	message         chan bool  // 节点间发信息的通道
	electCh         chan bool  // 选举的通道
	heartBeat       chan bool  // 心跳信号的通道
	heartbeatRe     chan bool  // 返回心跳信号的通道
	timeout         int        // 超时时间
}

// 0 表示还没上任， -1 没有 leader 则说明没有编号
var leader = Leader{0, -1}

func main() {
	// 过程，有三个节点，最初都是 follower
	// 若有 candidate 状态，进行投票和拉票
	// 会产生 leader

	// 创建3个节点
	for i := 0; i < raftCount; i++ {
		// 创建三个raft节点
		MakeNote(i)
	}
	// 加入服务端监听
	rpc.Register(new(Raft))
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	for {;
	}
}

// MakeNote 创建节点方法
func MakeNote(me int) *Raft {
	rf := &Raft{}
	rf.me = me
	// -1 代表谁都不投票，此时节点刚创建
	rf.votedFor = -1
	// o follower
	rf.state = 0
	rf.timeout = 0
	rf.currentLeader = -1
	// 节点任期
	rf.setTerm(0)
	/* 初始化相关通道
	message         chan bool // 初始化节点间发信息的通道
	electCh         chan bool // 初始化选举的通道
	heartBeat       chan bool // 初始化心跳信号的通道
	heartbeatRe     chan bool // 初始化返回心跳信号的通道
	*/
	rf.message = make(chan bool)
	rf.electCh = make(chan bool)
	rf.heartBeat = make(chan bool)
	rf.heartbeatRe = make(chan bool)
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 选举的协程
	go rf.election()
	// 心跳检查协程
	go rf.sendLeaderHeartBeat()
	return rf
}

// setTerm 设置节点任期
func (r *Raft) setTerm(term int) {
	r.currentTerm = term
}

// 选举方法
func (r *Raft) election() {
	// 设置标记，判断是否选出了 leader
	var result bool
	for {
		// 设置超时,150 到 300 随机数
		timeout := randRange(150, 300)
		r.lastMessageTime = millisecond()
		select {
		// 延迟等待一毫秒
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			fmt.Println("当前节点状态为：", r.state)
		}
		result = false
		for !result {
			// 选主逻辑
			result = r.electionOneRand(&leader)
		}
	}
}

// 随机值方法
func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// 获取当前时间，发送最后一条数据的时间
func millisecond() int64 {
	return time.Now().Unix() / int64(time.Millisecond)
}

// 实现选主逻辑
func (r *Raft) electionOneRand(leader *Leader) bool {
	// 定义超时
	var timeout int64
	timeout = 100
	// 投票数量记录
	var vote int
	// 是否开始心跳信号
	var triggerHeartbeat bool
	// 时间
	last := millisecond()
	// 用于返回值
	success := false
	// 给当前节点编程 candidate
	r.mu.Lock()
	// 修改转换状态
	r.becomeCandidate()
	r.mu.Unlock()
	fmt.Println("start electing leader")
	for {
		// 遍历所有节点拉选表
		for i := 0; i < raftCount; i++ {
			if i != r.me {
				// 拉选票
				go func() {
					if leader.LeaderId < 0 {
						// 设置投票
						r.electCh <- true
					}
				}()
			}
		}
		// 设置投票数量
		vote = 1
		// 遍历
		for i := 0; i < raftCount; i++ {
			// 计算投票数量
			select {
			case ok := <-r.electCh:
				if ok {
					// 投票数量加一
					vote++
					// 若选票个数 / 2 ，则成功当选 leader
					success = vote > raftCount/2
					if success && !triggerHeartbeat {
						// 更改状态变化成主节点，选主成功
						// 开始触发心跳检测
						triggerHeartbeat = true
						r.mu.Lock()
						// 更改状态为主
						r.becomeLeader()
						r.mu.Unlock()
						// 此时由  leader 向其他节点发送心跳信号
						r.heartBeat <- true
						fmt.Println(r.me, "号节点称为了 leader")
						fmt.Println("leader 开始发送心跳信号了")
					}
				}
			}
		}
		// 做最后的校验工作
		// 若不超时，且票数大于一半，则选举成功, break
		if timeout+last < millisecond() || (vote > raftCount/2 || r.currentLeader > -1) {
			break
		} else {
			// 等待操作
			select {
			case <-time.After(time.Duration(10) * time.Millisecond):
			}
		}
	}
	return success
}

// 修改状态 candidate
func (r *Raft) becomeCandidate() {
	r.state = 1
	r.setTerm(r.currentTerm + 1)
	r.votedFor = r.me
	r.currentLeader = -1
}

// 修改状态 leader
func (r *Raft) becomeLeader() {
	r.state = 2
	r.currentLeader = r.me
}

// sendLeaderHeartBeat 检查心跳,leader 节点发送心跳信号
// 顺便完成数据同步，先不实现
// 看 follower 是否存活
func (r *Raft) sendLeaderHeartBeat() {
	// 死循环
	for {
		select {
		case <-r.heartBeat:
			r.sendAppendEntriesImp1()
		}
	}
}

// 用于返回给 leader 的确认信号
func (r *Raft) sendAppendEntriesImp1() {
	// 如果是主就无需发送确认信号
	if r.currentLeader == r.me { // 此时是 leader
		// 记录确认信号的节点个数
		var successCount = 0
		// 设置确认信号
		for i := 0; i < raftCount; i++ {
			if i != r.me {
				go func() {
					//r.heartbeatRe <- true
					// RPC 分布式实现
					// 这里实际上相当于客户端
					rp, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
					if err != nil {
						log.Fatal(err)
					}
					// 接收服务器返回的信息
					// 接受服务端返回信息的变量
					var ok = false
					err = rp.Call("Raft.Communication", Param{"hello"}, &ok)
					if err != nil {
						log.Fatal(err)
					}
					if ok {
						r.heartbeatRe <- true
					}
				}()
			}
		}
		// 计算返回确认信号个数
		for i := 0; i < raftCount; i++ {
			select {
			case ok := <- r.heartbeatRe:
				if ok {
					successCount++
					if successCount > raftCount/2 {
						fmt.Println("投票选举成功，心跳信号OK")
						log.Fatal("程序结束")
					}
				}
			}
		}
	}
}

// RPC 规范，首字母大写
// 分布式通信

type Param struct {
	Msg string
}

// Communication 通信方法
func (r *Raft)Communication(p Param, a *bool) error {
	fmt.Println(p.Msg)
	*a = true
	return nil
}