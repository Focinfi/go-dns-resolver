package resolver

import (
	"fmt"
	"github.com/miekg/dns"
	"time"
)

type QueryType uint16

const (
	TypeA     = QueryType(dns.TypeA)
	TypeNS    = QueryType(dns.TypeNS)
	TypeMX    = QueryType(dns.TypeMX)
	TypeSOA   = QueryType(dns.TypeSOA)
	TypeCNAME = QueryType(dns.TypeCNAME)
	TypeTXT   = QueryType(dns.TypeTXT)
)

func (q QueryType) String() (queryTypeS string) {
	switch q {
	case QueryType(dns.TypeA):
		queryTypeS = "A"
	case QueryType(dns.TypeNS):
		queryTypeS = "NS"
	case QueryType(dns.TypeMX):
		queryTypeS = "MX"
	case QueryType(dns.TypeSOA):
		queryTypeS = "SOA"
	case QueryType(dns.TypeCNAME):
		queryTypeS = "CNAME"
	case QueryType(dns.TypeTXT):
		queryTypeS = "TXT"
	default:
		queryTypeS = "Unsportted Type"
	}
	return
}

type Result struct {
	Record   string
	Type     QueryType
	TTL      time.Duration
	Priority uint
	Content  string
}

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
