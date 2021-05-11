package taillog

import (
	"fmt"
	"github.com/GoWholeStack/oldboyGo/logAgent/etcd"
	"time"
)

// tailLogMgr tail 任务管理者
type tailLogMgr struct {
	logEntry []*etcd.LogEntry
	tailObjMap map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

var tailObjMgr *tailLogMgr

// Init 3.1 循环每一个收集项，创建一个 TailObjMgr
func Init(logEntryConf []*etcd.LogEntry) {
	tailObjMgr = &tailLogMgr{
		logEntry: logEntryConf,  // 将当前的日志收集项保存起来
		tailObjMap: make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), // 无缓冲区的channel
	}
	for _, logEntry := range logEntryConf {
		// 创造一个 tailTask 任务执行打开日志文件及发往 kafka
		// 初始化的时候起了多少个 tailObjMgr 都要记下来，为后续判断方便
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tailObjMgr.tailObjMap[mk] = tailObj
	}
	// 初始化完之后开始监听配置的变化
	go tailObjMgr.run()
}

// run 监听自己的 newConfChan 有了新的配置过来之久就做对应的处理(配置新增、配置删除、配置变更）
func (t *tailLogMgr)run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tailObjMap[mk]
				if ok {
					// 原来有，则不需要操作
					continue
				} else {
					// 原来没有，则表示新增项
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.tailObjMap[mk] = tailObj
				}
			}
			// 找出原 t.logEntry 中有，但是 newConf 中没有的则删除
			for _, c1 := range t.logEntry { // 从原配置中依次拿出原配置项
				isDelete := true
				for _, c2 := range newConf { // 从新匹配置中依次拿出新配置
					// 对原配置的内容和新配置的内容做比较，如果都存在则不做任何操作
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					// 将 c1 对应的 tailObj 停止
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					t.tailObjMap[mk].cancelFunc()
				}
			}
			// 2. 配置删除
			fmt.Println("新的配置来了：", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// NewConfChan 向外暴露 tailObjMgr 中的 newConfChan 通道
func NewConfChan() chan <- []*etcd.LogEntry {
	return tailObjMgr.newConfChan
}