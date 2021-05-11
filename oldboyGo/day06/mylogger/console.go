package mylogger

import (
	"fmt"
	"os"
	"time"
)

// 往终端写日志相关内容
// Logger 日志结构体
type ConseLogger struct {
	Level logLevel
}

// NewConseLog 日志对象构造方法
func NewConseLog(msg string) *ConseLogger {
	level, err := parseLogLevel(msg)
	if err != nil {
		panic(err)
	}
	return &ConseLogger{
		Level: level,
	}
}

// 根据日志级别控制日志的打印输出级别
func (c *ConseLogger) enable(loglevel logLevel) bool {
	// 表示输出用户传入的日志级别及该级别往后的日志级别信息，以实现开关功能
	return loglevel >= c.Level
}

// 日志记录函数方法
func (c *ConseLogger) log(format string, lv logLevel, agr ...interface{}) {
	if c.enable(lv) {
		now := time.Now()
		msg := fmt.Sprintf(format, agr...)
		timeFormat := now.Format("2006-01-02 15:04:05.000")
		// 获取出错内容的文件名、函数名、行号
		funcName, fileName, lineNo := getInfo(3)
		logLevelString := getLogString(lv)
		// 记录日志，os.Stdout 表示往终端输出
		fmt.Fprintf(os.Stdout, "%s <%s> [%s:%s:%d] %s\n", timeFormat, logLevelString, fileName, funcName, lineNo, msg)
	}
}

func (c *ConseLogger) Debug(format string, agr ...interface{}) {
	c.log(format, DEBUG, agr...)
}

func (c *ConseLogger) Tarce(format string, arg ...interface{}) {
	c.log(format, TARCE, arg...)
}

func (c *ConseLogger) Info(format string, arg ...interface{}) {
	c.log(format, INFO, arg...)
}

func (c *ConseLogger) Warning(format string, arg ...interface{}) {
	c.log(format, WARNING, arg...)
}

func (c *ConseLogger) Error(format string, agr ...interface{}) {
	c.log(format, ERROR, agr...)
}

func (c *ConseLogger) Fatal(format string, agr ...interface{}) {
	c.log(format, FATAL, agr...)
}
