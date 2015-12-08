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
}

func NewResolver(target, server string, otps map[string]interface{}) *Resolver {
	resolver := Resolver{Target: target, Server: server}
	resolver.client = &dns.Client{}
	resolver.msg = &dns.Msg{}

	// set DialTimeout
	if timeoutS := otps["timeout"]; timeoutS != nil {
		if timeoutI, ok := timeoutS.(int64); ok && timeoutI > 0 {
			resolver.client.DialTimeout = time.Duration(timeoutI)
		}
	}

	return &resolver
}

func (resolver *Resolver) Lookup(queryType QueryType) (results string, errors error) {
	res := make(chan string, 1)
	err := make(chan error, 1)
	resolver.Exchange(queryType, res, err)

	select {
	case results = <-res:
	case errors = <-err:
	}
	return
}

func (resolver Resolver) Exchange(queryType QueryType, results chan string, errors chan error) {
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
