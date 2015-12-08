package resolver

import (
	"testing"
	"time"
)

var resolver = NewResolver("focinfi.wang", "119.29.29.29", Config{Timeout: time.Second, RetryTimes: uint(3)})

func TestResovler(t *testing.T) {
	if res, err := resolver.Lookup(TypeNS); err == nil {
		res1 := res[0]
		t.Log(res1)
	} else {
		t.Error(err)
	}
}
