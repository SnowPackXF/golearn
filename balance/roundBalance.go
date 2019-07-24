package balance

import "errors"

func init()  {
	registBalance := RegistBalance{}
	registBalance.Regist("Round",&RoundBanance{})
}

type RoundBanance struct {
	indexInstance int
}

func (p *RoundBanance) Dobalance(instanceArr []*Instance,key ...string) (instance *Instance, err error) {
	if len(instanceArr) == 0 {
		err = errors.New("Instance is null!")
	}
	lens := len(instanceArr)
	p.indexInstance = (p.indexInstance + 1) % lens
	instance = instanceArr[p.indexInstance]
	return
}
