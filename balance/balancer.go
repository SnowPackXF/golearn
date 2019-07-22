package balance

type Balancer interface {
	Dobalance(instanceArr []*Instance) *Instance
}
