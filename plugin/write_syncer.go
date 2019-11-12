package plugin

import (
    "io"
)

//日志输出接口
type WriteSyncer interface {
    io.Writer
    Sync() error
}
