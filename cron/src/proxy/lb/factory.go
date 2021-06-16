package lb

type LoadBalanceType int

const (
	Random LoadBalanceType = iota
	RoundRobin
	WeightRound
	ConsistentHash
)

func GetLoadBalancer(loadBalanceType LoadBalanceType) LoadBalancer {
	switch loadBalanceType {
	case Random:
		return &RandomBalancer{}
	case RoundRobin:
		return &RoundRobinBalancer{}
	case WeightRound:
		return &WeightRoundRobinBalancer{}
	case ConsistentHash:
		return NewConsistentHashBalancer(10, nil)
	default:
		return &RandomBalancer{}
	}
}