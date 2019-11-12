package plugin

import (
    "os"
    "path/filepath"
    "time"

    rotatelogs "github.com/lestrrat-go/file-rotatelogs"

    "github.com/shmilwdc/glog/conf"
)

func NewTimeRollingFile(path, service string, maxAge int, rotationTime conf.TimeUnit) WriteSyncer {
    err := os.MkdirAll(path, 0766)
    if err != nil {
        panic(err)
    }

    rotateLogs, err := rotatelogs.New(
        filepath.Join(path, service+rotationTime.Format()+".log"),
        rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour),
        rotatelogs.WithRotationTime(rotationTime.Duration()),
    )
    if err != nil {
        panic(err)
    }

    return newLestrratWriteSyncer(rotateLogs)
}

type lestrratWriteSyncer struct {
    *rotatelogs.RotateLogs
}

func newLestrratWriteSyncer(l *rotatelogs.RotateLogs) *lestrratWriteSyncer {
    ws := &lestrratWriteSyncer{
        RotateLogs: l,
    }
    return ws
}

func (l *lestrratWriteSyncer) Sync() error {
    return nil
}
