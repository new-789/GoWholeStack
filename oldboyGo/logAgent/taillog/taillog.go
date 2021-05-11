package taillog

import (
	"context"
	"fmt"
	"github.com/GoWholeStack/oldboyGo/logAgent/kafka"
	"github.com/hpcloud/tail"
	//"time"
)

// 专门收集日志模块

// TailTask 一个日志收集的实例，包含要收集的日志路径和存储在 kafka 中的 topic 和 tail 对象
type TailTask struct {
	path string
	topic string
	instance *tail.Tail
	// 未来能够实现退出 t.run()
	ctx context.Context
	cancelFunc context.CancelFunc
}

// NewTailTask 构造函数，返回一个 TailObj 对象
func NewTailTask(path, topic string) (tailObj *TailTask)  {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path :path,
		topic: topic,
		ctx: ctx,
		cancelFunc: cancel,
	}
	// 根据路径去打开对应的日志文件
	tailObj.init()
	return
}

// Init 初始化连接 tail 并打开日志文件
func (t *TailTask)init() {
	// logEntry: 要收集的日志文件路径
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的那个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,                                 //
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Printf("tail file failed, err:%v\n", err)
	}
	// 开启 goroutine 采集日志并发送到 kafka 中，当 goroutine 执行的函数退出时 goroutine 则结束
	go t.run()
}

func (t *TailTask)run() {
	// 从 tailObj 通道中一行行的读取数据
	for {
		select {
		case <- t.ctx.Done():
			fmt.Printf("tail task:%s_%s end.....\n", t.path, t.topic)
			return
		case line := <- t.instance.Lines:
			// 3.2 发往 kafka
			//kafka.SendToKafka(t.topic, line.Text) // 函数调函数

			// 先把日志发到一个通道中，
			kafka.SendToChan(t.topic, line.Text)
			// 在 kafka 那个包中开启单独的 goroutine 去取日志数据并发送到 kafka
		}
	}
}

/*
// ReadChan 返回用来存储数据的 tail channel
func (t *TailTask)ReadChan() <-chan *tail.Line {
	return t.instance.Lines
}
 */