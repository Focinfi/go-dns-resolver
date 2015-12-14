package resolver

import (
	"github.com/miekg/dns"
)

type QueryType uint16

const (
	TypeA     = QueryType(dns.TypeA)
	TypeAAAA  = QueryType(dns.TypeAAAA)
	TypeNS    = QueryType(dns.TypeNS)
	TypeMX    = QueryType(dns.TypeMX)
	TypeSOA   = QueryType(dns.TypeSOA)
	TypeCNAME = QueryType(dns.TypeCNAME)
	TypeTXT   = QueryType(dns.TypeTXT)
	TypePTR   = QueryType(dns.TypePTR)
)

func (q QueryType) String() (queryTypeS string) {
	switch q {
	case TypeA:
		queryTypeS = "A"
	case TypeAAAA:
		queryTypeS = "AAAA"
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
	case TypePTR:
		queryTypeS = "PTR"
	default:
		queryTypeS = "Unsportted Type"
	}
	return
}
