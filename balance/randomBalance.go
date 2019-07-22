package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
}

func (p *RandomBalance) Dobalance(instanceArr []*Instance) (instance *Instance, error error) {

	i := len(instanceArr)
	if i == 0 {
		error = errors.New("InstanceArr is null")
		return
	}
	index := rand.Intn(i)
	return instanceArr[index], nil
}
