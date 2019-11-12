package plugin

import (
    "os"
    "path/filepath"

    "gopkg.in/natefinch/lumberjack.v2"
)

func NewSizeRollingFile(path, service string, maxSize, maxAge int, isCompress bool) WriteSyncer {
    err := os.MkdirAll(path, 0766)
    if err != nil {
        panic(err)
    }

    logger := &lumberjack.Logger{
        Filename:  filepath.Join(path, service+".log"),
        MaxSize:   maxSize,
        MaxAge:    maxAge,
        LocalTime: true,
        Compress:  isCompress,
    }

    return newLumberjackWriteSyncer(logger)
}

type lumberjackWriteSyncer struct {
    *lumberjack.Logger
}

func newLumberjackWriteSyncer(l *lumberjack.Logger) *lumberjackWriteSyncer {
    ws := &lumberjackWriteSyncer{
        Logger: l,
    }
    return ws
}

func (l *lumberjackWriteSyncer) Sync() error {
    return nil
}
