package balance

import "errors"

type RoundBanance struct {
	indexInstance int
}

func (p *RoundBanance) Dobalance(instanceArr []*Instance) (instance *Instance, err error) {
	if len(instanceArr) == 0 {
		err = errors.New("Instance is null!")
	}
	lens := len(instanceArr)
	p.indexInstance = (p.indexInstance + 1) % lens
	return
}
