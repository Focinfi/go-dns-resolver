package resolver

import (
	"github.com/miekg/dns"
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
