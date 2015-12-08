package resolver

import (
	"time"
)

type Config struct {
	Timeout    time.Duration
	RetryTimes uint
}
