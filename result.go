package resolver

import (
	"time"
)

type Result struct {
	Record   string
	Type     QueryType
	TTL      time.Duration
	Priority uint
	Content  string
}
