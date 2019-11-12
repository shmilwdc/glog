package conf

import (
    "time"
)

type TimeUnit string

const (
    Minute = "minute"
    Hour   = "hour"
    Day    = "day"
    Month  = "month"
)

func (t TimeUnit) Format() string {
    switch t {
    case Minute:
        return ".%Y%m%d%H%M"
    case Hour:
        return ".%Y%m%d%H"
    case Day:
        return ".%Y%m%d"
    case Month:
        return ".%Y%m"
    default:
        return ".%Y%m%d"
    }
}

func (t TimeUnit) Duration() time.Duration {
    switch t {
    case Minute:
        return time.Minute
    case Hour:
        return time.Hour
    case Day:
        return time.Hour * 24
    case Month:
        return time.Hour * 24 * 30
    default:
        return time.Hour * 24
    }
}
