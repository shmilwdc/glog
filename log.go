package log

import (
    "github.com/shmilwdc/glog/core"
)

type Logger interface {
    //普通日志,如果有args，需要格式化
    Debug(string, ...interface{})
    Info(string, ...interface{})
    Warn(string, ...interface{})
    Error(string, ...interface{})
    Panic(string, ...interface{})
    Fatal(string, ...interface{})
    //需要格式化日志 ，最后一个是context
    Debugf(string, ...interface{})
    Infof(string, ...interface{})
    Warnf(string, ...interface{})
    Errorf(string, ...interface{})
    Panicf(string, ...interface{})
    Fatalf(string, ...interface{})
    SetLevel(string)
}

var std Logger = core.New()

func SetLogger(logger Logger) {
    std = logger
}

//普通日志
func Debug(msg string, args ...interface{}) {
    std.Debug(msg, args...)
}
func Info(msg string, args ...interface{}) {
    std.Info(msg, args...)
}
func Warn(msg string, args ...interface{}) {
    std.Warn(msg, args...)
}
func Error(msg string, args ...interface{}) {
    std.Error(msg, args...)
}
func Panic(msg string, args ...interface{}) {
    std.Panic(msg, args...)
}
func Fatal(msg string, args ...interface{}) {
    std.Fatal(msg, args...)
}

//其他日志 如：HTTP RPC日志
func Debugf(format string, args ...interface{}) {
    std.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
    std.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) {
    std.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
    std.Errorf(format, args...)
}
func Panicf(format string, args ...interface{}) {
    std.Panicf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
    std.Fatalf(format, args...)
}

func SetLevel(level string) {
    std.SetLevel(level)
}
