# go-dns-resovler

DNS resolver in Golang, based on [miekg/dns](github.com/miekg/dns).

### Goal

1. Simple interface
2. Rich and neat output
3. Easy to figure
4. High performance

### Example

```go
package main

import (
  dns "github.com/Focinfi/go-dns-resovler"
  "log"
)

func main() {
  target := "f.focinfi.wang"
  server := "119.29.29.29"
  config := map[string]interface{}{}
  config["timeout"] = uint(1)

  resolver := dns.NewResolver(target, server, config)
  if res, err := resolver.Lookup(dns.TypeCNAME); err == nil {
    log.Println(res)
  } else {
    log.Fatal(err)
  }
}
```

