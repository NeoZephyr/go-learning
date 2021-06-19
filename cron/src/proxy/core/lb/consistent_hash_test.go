package lb

import (
	"fmt"
	"testing"
)

func TestConsistentHashBalance(t *testing.T) {
	rb := NewConsistentHashBalancer(10, nil)
	rb.Add("127.0.0.1:2002")
	rb.Add("127.0.0.1:2003")
	rb.Add("127.0.0.1:2004")
	rb.Add("127.0.0.1:2005")
	rb.Add("127.0.0.1:2006")

	fmt.Println(rb.Get("a"))
	fmt.Println(rb.Get("b"))
	fmt.Println(rb.Get("c"))
	fmt.Println(rb.Get("a"))
}