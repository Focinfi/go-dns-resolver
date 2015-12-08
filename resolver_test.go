package resolver

import (
	"testing"
	"time"
)

func TestResovler(t *testing.T) {
	target := "f.focinfi.wang"
	server := "119.29.29.29"
	config := Config{Timeout: uint(time.Second), RetryTimes: uint(3)}
	resolver := NewResolver(target, server, config)

	if res, err := resolver.Lookup(TypeCNAME); err == nil {
		t.Log(res)
	} else {
		t.Error(err)
	}
}
