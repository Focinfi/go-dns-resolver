package resolver

import (
	"testing"
)

func TestResovler(t *testing.T) {
	target := "f.focinfi.wang"
	server := "119.29.29.29"
	resolver := NewResolver(target, server, nil)

	if res, err := resolver.Lookup(TypeCNAME); err == nil {
		t.Log(res)
	} else {
		t.Error(err)
	}
}
