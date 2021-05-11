package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// 自定义日之库

// 自定义类型为 unit16 的别名
type logLevel uint16

// Logger 接口
type Logger interface {
	Debug(format string, agr ...interface{})
	Tarce(format string, agr ...interface{})
	Info(format string, agr ...interface{})
	Warning(format string, agr ...interface{})
	Error(format string, agr ...interface{})
	Fatal(format string, agr ...interface{})
}

// 自定义日志级别
const (
	UNKNOWN logLevel = iota
	DEBUG
	TARCE
	INFO
	WARNING
	ERROR
	FATAL
)

// 判断日志级别
func parseLogLevel(s string) (logLevel, error) {
	s = strings.ToLower(s) // 将字符串转换为小写
	switch s {
	case "debug":
		return DEBUG, nil
	case "tarce":
		return TARCE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("错误！无效的日志级别")
		return UNKNOWN, err
	}
}

// 将自定义的日志级别类型转换为字符串类型返回以便写入文件
func getLogString(lv logLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TARCE:
		return "TARCE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}

// 获取调用 runtime.Caller() 方法的函数名、文件名以及行号
func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName = strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]
	fileName = path.Base(file) // 获取一个路径的最后一个内容
	lineNo = line
	return
}
