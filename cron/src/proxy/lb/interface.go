package lb

type LoadBalancer interface {
	Add(params ...string) error
	Get(string) (string, error)
}
