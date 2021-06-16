package load_balance

import "testing"

func TestRoundRobinBalance(t *testing.T) {
	rb := &RoundRobinBalance{}
	rb.Add("127.0.0.1:2002")
	rb.Add("127.0.0.1:2003")
	rb.Add("127.0.0.1:2004")
	rb.Add("127.0.0.1:2005")
	rb.Add("127.0.0.1:2006")

	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
	println(rb.Next())
}
