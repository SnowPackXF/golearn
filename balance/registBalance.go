package balance

var rb = &RegistBalance{
	RegistMap: map[string] Balancer{},
}
type RegistBalance struct {
	RegistMap map[string] Balancer
}

func (p *RegistBalance)Regist(name string,balancer Balancer)  {
	rb.RegistMap[name] = balancer
}

func (p *RegistBalance)Dobalance(instanceArr []*Instance,key ...string) (instance *Instance,err error)  {

	strategy,_:=rb.RegistMap[key[0]]
	instance,err = strategy.Dobalance(instanceArr)
	return
}