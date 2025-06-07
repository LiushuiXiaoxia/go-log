package main

import (
	"C"
	"fmt"
	"os"
	"sync"
	"time"
)

// Level 定义日志级别
type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

type Logger struct {
	mu       sync.Mutex
	file     *os.File
	filePath string
}

// NewLogger 创建日志对象，传入日志文件目录
func NewLogger(dir string) (*Logger, error) {
	logger := &Logger{}
	err := logger.rotateFile(dir)
	if err != nil {
		return nil, err
	}
	return logger, nil
}

// rotateFile 每天生成一个新的日志文件，格式 log-YYYY-MM-DD.txt
func (l *Logger) rotateFile(dir string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file != nil {
		l.file.Close()
	}

	filename := fmt.Sprintf("%s/golog-%s.txt", dir, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	l.file = file
	l.filePath = filename
	return nil
}

// log 写日志的核心函数
func (l *Logger) log(level Level, tag, msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now().Format("2006-01-02 15:04:05")
	logLine := fmt.Sprintf("%s [%s] [%s] %s\n", now, level.String(), tag, msg)
	if l.file != nil {
		l.file.WriteString(logLine)
	}
}

// 对外的日志接口

func (l *Logger) Debug(tag, msg string) {
	l.log(DEBUG, tag, msg)
}

func (l *Logger) Info(tag, msg string) {
	l.log(INFO, tag, msg)
}

func (l *Logger) Warn(tag, msg string) {
	l.log(WARN, tag, msg)
}

func (l *Logger) Error(tag, msg string) {
	l.log(ERROR, tag, msg)
}

// Close 关闭日志文件
func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

var logger *Logger

//export InitLogger
func InitLogger(logDir *C.char) C.int {
	dir := C.GoString(logDir)

	filename := fmt.Sprintf("%s/log-%s.txt", dir, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return -1
	}
	logger = &Logger{
		file:     file,
		filePath: filename,
	}
	return 0
}

// helper to convert C strings safely
func cStrToGo(cstr *C.char) string {
	if cstr == nil {
		return ""
	}
	return C.GoString(cstr)
}

//export LogDebug
func LogDebug(tag *C.char, msg *C.char) {
	if logger == nil {
		return
	}
	logger.log(DEBUG, cStrToGo(tag), cStrToGo(msg))
}

//export LogInfo
func LogInfo(tag *C.char, msg *C.char) {
	if logger == nil {
		return
	}
	logger.log(INFO, cStrToGo(tag), cStrToGo(msg))
}

//export LogWarn
func LogWarn(tag *C.char, msg *C.char) {
	if logger == nil {
		return
	}
	logger.log(WARN, cStrToGo(tag), cStrToGo(msg))
}

//export LogError
func LogError(tag *C.char, msg *C.char) {
	if logger == nil {
		return
	}
	logger.log(ERROR, cStrToGo(tag), cStrToGo(msg))
}

//export CloseLogger
func CloseLogger() {
	if logger != nil && logger.file != nil {
		logger.file.Close()
		logger = nil
	}
}

func main() {}
