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
	case TypeA:
		queryTypeS = "A"
	case TypeNS:
		queryTypeS = "NS"
	case TypeMX:
		queryTypeS = "MX"
	case TypeSOA:
		queryTypeS = "SOA"
	case TypeCNAME:
		queryTypeS = "CNAME"
	case TypeTXT:
		queryTypeS = "TXT"
	default:
		queryTypeS = "Unsportted Type"
	}
	return
}
