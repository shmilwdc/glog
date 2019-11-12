package core

import (
    "os"
    "strings"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"

    "github.com/shmilwdc/glog/conf"
    "github.com/shmilwdc/glog/plugin"
)

type Log struct {
    logger *zap.Logger
    level  zap.AtomicLevel
}

func parseLevel(l string) zapcore.Level {
    switch strings.ToLower(l) {
    case "debug":
        return zapcore.DebugLevel
    case "info":
        return zapcore.InfoLevel
    case "warn", "warning":
        return zapcore.WarnLevel
    case "error":
        return zapcore.ErrorLevel
    case "panic", "dpanic":
        return zapcore.PanicLevel
    case "fatal":
        return zapcore.FatalLevel
    default:
        return zapcore.DebugLevel
    }
}

func New(opts ...conf.Option) *Log {
    o := &conf.DefaultConf
    for _, opt := range opts {
        opt(o)
    }

    var writers []zapcore.WriteSyncer
    if o.IsConsole {
        writers = append(writers, os.Stdout)
    }
    switch o.RotateType {
    case conf.Size:
        writers = append(writers, zapcore.AddSync(plugin.NewSizeRollingFile(o.LogPath, o.LogName, o.MaxSize, o.MaxAge, o.IsCompress)))
    case conf.Time:
        writers = append(writers, zapcore.AddSync(plugin.NewTimeRollingFile(o.LogPath, o.LogName, o.MaxAge, o.TimeUnit)))
    }

    atomicLevel := zap.NewAtomicLevel()
    atomicLevel.SetLevel(parseLevel(o.LogLevel))

    logger := newZapLogger(atomicLevel, parseLevel(o.Stacktrace), zapcore.NewMultiWriteSyncer(writers...))
    zap.RedirectStdLog(logger)

    logger = logger.With(zap.String("service", o.LogName))

    return &Log{logger: logger, level: atomicLevel}
}

func newZapLogger(level zap.AtomicLevel, stacktrace zapcore.Level, ws zapcore.WriteSyncer) *zap.Logger {
    encoderConfig := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
        EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
        EncodeDuration: zapcore.SecondsDurationEncoder, //
        EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
        EncodeName:     zapcore.FullNameEncoder,
    }

    encoder := zapcore.NewJSONEncoder(encoderConfig)

    core := zapcore.NewCore(encoder, ws, level)

    return zap.New(core, zap.AddCaller(), zap.AddStacktrace(stacktrace), zap.AddCallerSkip(2))
}
