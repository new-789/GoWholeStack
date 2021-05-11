package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件中写入日志相关代码

var (
	// MaxSize 日志通道大小
	MaxSize int = 50000
)

//  FileLogger 文件日志结构体
type FileLogger struct {
	Level       logLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMag
}

type logMag struct {
	level     logLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, filePath, fileName string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxSize,
		logChan:     make(chan *logMag, MaxSize),
	}
	err = fl.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

// 根据指定的日志文件路径和文件名打开日志文件
func (f *FileLogger) initFile() error {
	FullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(FullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return err
	}
	errFileObj, err := os.OpenFile(FullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("open errfile failed, err:", err)
		return err
	}
	// 日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	// 开启一个后台 goroutine 执行，往文件中写日志内容操作
	go f.writeLogBackground()
	return nil
}

// 根据日志级别控制日志的打印输出级别，即按段是否需要记录该日志
func (f *FileLogger) enable(loglevel logLevel) bool {
	// 表示输出用户传入的日志级别及该级别往后的日志级别信息，以实现开关功能
	return loglevel >= f.Level
}

// 获取文件大小方法，判断文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file size failed, err:", err)
		return false
	}
	// 如果当前文件大小大于等于设定的日志文件最大值，则应该返回 true，表示需要切割文件
	return fileInfo.Size() >= f.maxFileSize
}

// 切割文件方法
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割文件
	// 2. rename 备份一下
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file info failed, err:", err)
		return nil, err
	}
	// 拿到当前的日志文件完整路径
	logName := path.Join(f.filePath, fileInfo.Name())
	newFileName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	// 1. 关闭日志文件
	file.Close()
	// 将现有日志文件进行备份
	os.Rename(logName, newFileName)
	// 3. 打开一个新的日志文件赋值给f.fileObj
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("open new log file failed, err:", err)
		return nil, err
	}
	return fileObj, nil
}

// writeLogBackground 后台往文件中写日志方法
func (f *FileLogger) writeLogBackground() {
	for {
		// 检查日志文件大小，并切割文件
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj) // 普通日志文件
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		// 往文件中记录日志信息
		select {
		case logTmp := <- f.logChan:
			// 拼出日志内容
			logInfo := fmt.Sprintf("%s <%s> [%s:%s:%d] %s\n", logTmp.timestamp, getLogString(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
			fmt.Fprintf(f.fileObj, logInfo)

			if logTmp.level >= ERROR {
				if f.checkSize(f.errFileObj) {
					newFile, err := f.splitFile(f.errFileObj)
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				// 如果要记录的日志大于等于 error 级别，还要在 err 日志文件中在记录一遍
				fmt.Fprintf(f.fileObj, logInfo)
			}
		default:
			// 若取不到日志先休息 500 毫秒
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// 日志记录函数方法,往日志中写入文件通过 goroutine 方法
func (f *FileLogger) log(format string, lv logLevel, agr ...interface{}) {
	if f.enable(lv) {
		now := time.Now()
		timeFormat := now.Format("2006-01-02 15:04:05.000")
		msg := fmt.Sprintf(format, agr...)
		// 获取出错内容的文件名、函数名、行号
		funcName, fileName, lineNo := getInfo(3)
		// 先把日志放到通道中
		// 1. 造一个 logMsg 对象
		logTmp := &logMag{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: timeFormat,
			line:      lineNo,
		}
		select {
		case f.logChan <- logTmp:
		default:
			// 执行到这里说明channel以满，此时扔掉日志保证业务代码顺利执行不出现阻塞
		}
	}
}

// Debug ...
func (f *FileLogger) Debug(format string, agr ...interface{}) {
	f.log(format, DEBUG, agr...)
}

// Tarce ...
func (f *FileLogger) Tarce(format string, arg ...interface{}) {
	f.log(format, TARCE, arg...)
}

// Info ...
func (f *FileLogger) Info(format string, arg ...interface{}) {
	f.log(format, INFO, arg...)
}

// Warning ...
func (f *FileLogger) Warning(format string, arg ...interface{}) {
	f.log(format, WARNING, arg...)
}

// Error ...
func (f *FileLogger) Error(format string, agr ...interface{}) {
	f.log(format, ERROR, agr...)
}

// Fatal ...
func (f *FileLogger) Fatal(format string, agr ...interface{}) {
	f.log(format, FATAL, agr...)
}

// Close ...
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
