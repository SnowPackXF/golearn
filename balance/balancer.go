package balance

type Balancer interface {

	Dobalance(instanceArr []*Instance,key ...string) (*Instance,error)
}
