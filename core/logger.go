package core

import (
	"fmt"
	"io"
	//"io/ioutil"
	"log"
	"os"
	//"path"
	"runtime/debug"
)

// Simple logger package to log to stdout

// Logger is the tmail log service
type Logger struct {
	debugEnabled bool
	DebugLogger  *log.Logger
	InfoLogger   *log.Logger
	err          *log.Logger
	trace        *log.Logger
}

func NewLogger(out io.Writer, debugEnabled bool) (*Logger, error) {
	hostname, _ := os.Hostname()
	return &Logger{
		debugEnabled: debugEnabled,
		DebugLogger:  log.New(out, "["+hostname+"] ", log.Ldate|log.Ltime|log.Lmicroseconds),
		InfoLogger:   log.New(out, "["+hostname+"] ", log.Ldate|log.Ltime|log.Lmicroseconds),
		err:          log.New(out, "["+hostname+"] ", log.Ldate|log.Ltime|log.Lmicroseconds),
		trace:        log.New(out, "["+hostname+"] ", log.Ldate|log.Ltime|log.Lshortfile),
	}, nil
}

func (l *Logger) Debug(v ...interface{}) {
	if !l.debugEnabled {
		return
	}
	msg := "DEBUG -"
	for i := range v {
		msg = fmt.Sprintf("%s %v", msg, v[i])
	}
	l.DebugLogger.Println(msg)
}

func (l *Logger) Info(v ...interface{}) {
	msg := "INFO -"
	for i := range v {
		msg = fmt.Sprintf("%s %v", msg, v[i])
	}
	l.InfoLogger.Println(msg)
}

func (l *Logger) Error(v ...interface{}) {
	msg := "ERROR -"
	for i := range v {
		msg = fmt.Sprintf("%s %v", msg, v[i])
	}
	l.err.Println(msg)
}

func (l *Logger) Trace(v ...interface{}) {
	stack := debug.Stack()
	msg := "TRACE -"
	for i := range v {
		msg = fmt.Sprintf("%s %v", msg, v[i])
	}
	msg += "\r\nStack: \r\n" + fmt.Sprintf("%s", stack)
	l.trace.Println(msg)
}

// nsq interface
func (l *Logger) Output(calldepth int, s string) error {
	l.Debug(s)
	return nil
}

// gorm interface
func (l *Logger) Print(v ...interface{}) {
	l.Debug(v)
}
