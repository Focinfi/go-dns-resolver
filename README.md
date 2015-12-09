# go-dns-resolver

DNS resolver in Golang, based on [miekg/dns](github.com/miekg/dns).

### Goal

1. Simple interface
2. Rich and neat output
3. Easy to figure
4. High performance

### Install
```shell
go get github.com/Focinfi/go-dns-resolver
```

### Example

```go
package main

import (
  dns "github.com/Focinfi/go-dns-resolver"
  "log"
)

func main() {
  domains := []string{"google.com", "twitter.com"}
  types := []dns.QueryType{dns.TypeA, dns.TypeNS, dns.TypeMX, dns.TypeTXT}

  // Set timeout
  dns.Config.SetTimeout(uint(2))

  // Create and setup resolver with domains and types
  resolver := dns.NewResolver("119.29.29.29")
  resolver.Targets(domains...).Types(types...)
  // Lookup
  res := resolver.Lookup()

  //res is a map[string]*ResultItem, key is the domain
  for target := range res.ResMap {
    log.Printf("%v: \n", target)
    for _, r := range res.ResMap[target] {
      if r != nil {
        log.Println(r.Record, r.Type, r.Ttl, r.Priority, r.Content)
      }
    }
  }
}

```

Output is:
``` shell
twitter.com: 
twitter.com A 10m3s 0 78.16.49.15
twitter.com NS 11h49m58s 0 ns1.p34.dynect.net.
twitter.com NS 11h49m58s 0 ns4.p34.dynect.net.
twitter.com NS 11h49m58s 0 ns3.p34.dynect.net.
twitter.com NS 11h49m58s 0 ns2.p34.dynect.net.
google.com: 
google.com TXT 19m26s 0 v=spf1 include:_spf.google.com ~all
google.com A 2m31s 0 216.58.197.110
google.com NS 7h40m6s 0 ns1.google.com.
google.com NS 7h40m6s 0 ns3.google.com.
google.com NS 7h40m6s 0 ns2.google.com.
google.com NS 7h40m6s 0 ns4.google.com.
google.com MX 10m0s 20 alt1.aspmx.l.google.com.
google.com MX 10m0s 10 aspmx.l.google.com.
google.com MX 10m0s 50 alt4.aspmx.l.google.com.
google.com MX 10m0s 40 alt3.aspmx.l.google.com.
google.com MX 10m0s 30 alt2.aspmx.l.google.com.
[Finished in 2.3s]
```

### Todo

1. Support Config.RetryTimes
2. Support more DNS record types

