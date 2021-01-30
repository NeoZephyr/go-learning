package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	var name string
	greeting, err := hello(name)

	if greeting != "" || err == nil {
		t.Errorf("Nonempty greeting and nil error, but it should not be. (name=%q)", name)
	}

	name = "Robert"
	greeting, err = hello(name)

	if greeting == "" || err != nil {
		t.Errorf("Empty greeting and nonNil error, but it should not be. (name=%q)", name)
	}
}

func TestFail(t *testing.T) {
	// t.Fail()
	t.FailNow()
	t.Log("Failed.")
}

// go test -bench=. -run=^$
// go test -bench=. -cpu=128 -run=^$
func BenchmarkGetPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getPrimes(1000)
	}
}
