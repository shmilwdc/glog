package conf

type Option func(*Conf)

func WithLogPath(logPath string) Option {
    return func(conf *Conf) {
        conf.LogPath = logPath
    }
}

func WithLogName(logName string) Option {
    return func(conf *Conf) {
        conf.LogName = logName
    }
}

func WithLogLevel(logLevel string) Option {
    return func(conf *Conf) {
        conf.LogLevel = logLevel
    }
}

func WithRotateType(rotateType RotateType) Option {
    return func(conf *Conf) {
        conf.RotateType = rotateType
    }
}

func WithMaxSize(maxSize int) Option {
    return func(conf *Conf) {
        conf.MaxSize = maxSize
    }
}

func WithMaxAge(maxAge int) Option {
    return func(conf *Conf) {
        conf.MaxAge = maxAge
    }
}

func WithTimeUnit(timeUnit TimeUnit) Option {
    return func(conf *Conf) {
        conf.TimeUnit = timeUnit
    }
}

func WithStacktrace(stacktrace string) Option {
    return func(conf *Conf) {
        conf.Stacktrace = stacktrace
    }
}

func WithIsCompress(isCompress bool) Option {
    return func(conf *Conf) {
        conf.IsCompress = isCompress
    }
}

func WithIsConsole(isConsole bool) Option {
    return func(conf *Conf) {
        conf.IsConsole = isConsole
    }
}
