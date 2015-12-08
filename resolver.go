package resolver

import (
	"fmt"
	"github.com/miekg/dns"
	"time"
)

type Resolver struct {
	Target string
	Server string
	client *dns.Client
	msg    *dns.Msg
	Config Config
}

func NewResolver(target, server string, config Config) *Resolver {
	resolver := Resolver{Target: target, Server: server, Config: config}
	resolver.client = &dns.Client{}
	resolver.msg = &dns.Msg{}
	// set DialTimeout
	resolver.client.DialTimeout = time.Duration(int64(config.Timeout))
	return &resolver
}

// Lookup querys the results of queryType via dns.Client
func (resolver *Resolver) Lookup(queryType QueryType) (results []*ResultItem, errors error) {
	resolver.msg.SetQuestion(resolver.Target+".", uint16(queryType))
	res, _, err := resolver.client.Exchange(resolver.msg, resolver.Server+":53")
	if err != nil {
		errors = err
	} else {
		ans := res.Answer
		if len(ans) == 0 {
			errors = fmt.Errorf("No %v result", queryType.String())
		} else {
			fmt.Println(len(res.Answer))
			for _, answer := range res.Answer {
				result := NewResultItemWithDnsRP(queryType, answer)
				result.Record = resolver.Target
				results = append(results, result)
			}
		}
	}
	return
}
