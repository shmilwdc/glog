package conf

type Conf struct {
    LogPath    string     //日志保存的路径
    LogName    string     //日志保存的名称
    LogLevel   string     //日志记录的级别
    RotateType RotateType //日志切割的类型 size/time
    MaxSize    int        //日志分割的尺寸 mb
    MaxAge     int        //分割日志的最大保存时间 day
    TimeUnit   TimeUnit   //分割日志的时间类型 hour/day
    Stacktrace string     //记录堆栈的级别
    IsCompress bool       //是否压缩
    IsConsole  bool       //是否标准输出console输出
}

var DefaultConf = Conf{
    LogPath:    "/var/log",
    LogName:    "output",
    LogLevel:   "debug",
    RotateType: Time,
    MaxSize:    100,
    MaxAge:     7,
    TimeUnit:   Hour,
    Stacktrace: "error",
    IsCompress: false,
    IsConsole:  true,
}
