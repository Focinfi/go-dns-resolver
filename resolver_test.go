package resolver

import (
	"testing"
)

func TestLookup(t *testing.T) {
	var resolver = NewResolver("119.29.29.29")
	Config.SetTimeout(uint(2))
	Config.RetryTimes = uint(4)

	resolver.Targets("www.google.com").Types(TypeA, TypeAAAA, TypeMX)
	res := resolver.Lookup()
	for target := range res.ResMap {
		t.Logf("%v: ", target)
		for _, r := range res.ResMap[target] {
			if r != nil {
				t.Log(r.Record, r.Type, r.Ttl, r.Priority, r.Content)
			}
		}
	}
}

func BenchmarkLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var resolver = NewResolver("119.29.29.29")
		resolver.Targets("youtube.com", "google.com", "twitter.com", "baidu.com").Types(TypeA, TypeAAAA, TypeMX, TypeTXT)
		resolver.Lookup()
	}
}
