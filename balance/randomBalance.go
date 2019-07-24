package balance

import (
	"errors"
	"math/rand"

)

func init()  {
	registBalance := new(RegistBalance)
	registBalance.Regist("Random", &RandomBalance{})
}
type RandomBalance struct {
}

func (p *RandomBalance) Dobalance(instanceArr []*Instance,key ...string) (instance *Instance, error error) {

	i := len(instanceArr)
	if i == 0 {
		error = errors.New("InstanceArr is null")
		return
	}
	index := rand.Intn(i)
	return instanceArr[index], nil
}
