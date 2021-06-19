package lb

import "testing"

func TestWeightRoundRobinBalance(t *testing.T) {
	rb := &WeightRoundRobinBalancer{}
	rb.Add("127.0.0.1:2002", "6")
	rb.Add("127.0.0.1:2003", "3")
	rb.Add("127.0.0.1:2004", "2")

	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
}