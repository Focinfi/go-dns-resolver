package resolver

import (
	"fmt"
	"github.com/miekg/dns"
)

type Resolver struct {
	Server string
	Query  *Query
}

func NewResolver(server string) *Resolver {
	resolver := Resolver{Server: server + ":53"}
	return &resolver
}

func (r *Resolver) Targets(targets ...string) *Query {
	query := NewQueryWithTargets(targets...)
	r.Query = query
	return query
}

func (r *Resolver) Lookup() *Result {
	result := Result{Server: r.Server, ResMap: map[string][]*ResultItem{}}

	resultsChan := make(chan []*ResultItem, r.Query.Count())
	for _, queryItem := range r.Query.Items {
		target := queryItem.Target
		for _, t := range queryItem.Types {
			go GoExchange(target, r.Server, t, resultsChan)
		}
	}

	for i := 0; i < r.Query.Count(); i++ {
		res := <-resultsChan
		if len(res) > 0 {
			target := res[0].Record
			result.ResMap[target] = append(result.ResMap[target], res...)
		}
	}
	return &result
}

func GoExchange(target string, server string, queryType QueryType, resultsChan chan []*ResultItem) {
	if results, err := Exchange(target, server, queryType); err == nil {
		resultsChan <- results
	} else {
		resultsChan <- []*ResultItem{}
	}
}

func Exchange(target string, server string, queryType QueryType) (results []*ResultItem, errors error) {
	results = []*ResultItem{}
	msg := &dns.Msg{}
	msg.SetQuestion(target+".", uint16(queryType))
	client := &dns.Client{DialTimeout: Config.Timeout}
	res, _, err := client.Exchange(msg, server)
	if err != nil {
		errors = err
	} else {
		ans := res.Answer
		if len(ans) == 0 {
			errors = fmt.Errorf("No %v result", queryType.String())
		} else {
			for _, answer := range res.Answer {
				result := NewResultItemWithDnsRP(queryType, answer)
				result.Record = target
				results = append(results, result)
			}
		}
	}
	return
}
