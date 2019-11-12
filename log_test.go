package log

import (
    "context"
    "testing"

    "github.com/shmilwdc/glog/conf"
    "github.com/shmilwdc/glog/core"
    "github.com/shmilwdc/glog/tracer"
)

func init() {
    SetLogger(core.New(
        conf.WithLogPath("/tmp"),
        conf.WithLogName("test"),
        conf.WithLogLevel("info"),
        conf.WithMaxSize(1),
        conf.WithRotateType("size"),
        conf.WithRotateType(""),
        conf.WithIsConsole(true),
    ))
}

func TestSetLogger(t *testing.T) {
    for i := 0; i < 1e3; i++ {
        Debug("debug test")
        Debug("debug test", context.Background())
        Infof("hello %s", "world", context.WithValue(context.Background(), tracer.LogTraceKey, "46b1506e7332f7c1:7f75737aa70629cc:3bb947500f42ad71"))
        Infof("hello %s %d", "world", i, context.Background())
    }
    SetLevel("debug")
    Debug("debug test")
}

func BenchmarkSetLogger(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            Error("error")
        }
    })
}