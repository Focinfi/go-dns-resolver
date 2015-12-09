package resolver

import (
	"time"
)

type Configuration struct {
	RetryTimes uint
}

var Config = Configuration{RetryTimes: uint(0)}

func SetTimeout(seconds uint) {
	Client.DialTimeout = time.Second * time.Duration(seconds)
}
