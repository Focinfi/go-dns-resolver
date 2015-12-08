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

you can get this package using `resolver`.

### Example

```go
package main

import (
  dns "github.com/Focinfi/go-dns-resolver"
  "log"
  "time"
)

func main() {
  target := "f.focinfi.wang"
  server := "119.29.29.29"
  // set Timeout is 1 second and retry 3 times if failed
  config := dns.Config{Timeout: time.Second, RetryTimes: uint(3)}
  resolver := dns.NewResolver(target, server, config)

  if res, err := resolver.Lookup(dns.TypeCNAME); err == nil {
    // res is a array of ResultItem
    for _, r := range res {
      log.Println(r.Record, r.Type, r.Ttl, r.Priority, r.Content)
    }
  } else {
    log.Fatalln(err)
  }
}
```

