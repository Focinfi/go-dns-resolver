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

func (resolver *Resolver) Lookup(queryType QueryType) (results string, errors error) {
	res := make(chan string, 1)
	err := make(chan error, 1)
	results, errors = resolver.GoExchange(queryType, res, err)
	// retry resolver.Config.RetryTimes if fails
	for i := 0; results == "" && i < int(resolver.Config.RetryTimes); i++ {
		results, errors = resolver.GoExchange(queryType, res, err)
	}
	return
}

// GoExchange run Exchange() in goroutine and wait for the results
func (resolver *Resolver) GoExchange(queryType QueryType, res chan string, err chan error) (results string, errors error) {
	go resolver.Exchange(queryType, res, err)
	select {
	case results = <-res:
	case errors = <-err:
	}
	return
}

// Exchage query the results for queryType via dns.Client
func (resolver *Resolver) Exchange(queryType QueryType, results chan string, errors chan error) {
	resolver.msg.SetQuestion(resolver.Target+".", uint16(queryType))
	res, _, err := resolver.client.Exchange(resolver.msg, resolver.Server+":53")
	if err != nil {
		errors <- err
	} else {
		ans := res.Answer
		if len(ans) == 0 {
			errors <- fmt.Errorf("No %v result", queryType.String())
		} else {
			results <- res.String()
		}
	}
}
