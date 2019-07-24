package main

import (
	"fmt"
	"golearn/day7/balance"
	"math/rand"
	"time"
)

func main() {
	var arr []*balance.Instance
	for i := 0; i < 10; i++ {
		newInstance := &balance.Instance{
			Ip:   fmt.Sprintf("192.31.%d,%d", rand.Intn(255), rand.Intn(255)),
			Port: 8888,
		}
		arr = append(arr, newInstance)
	}
	regist := balance.RegistBalance{}
	for {
		instance,_ := regist.Dobalance(arr,"Round")
		fmt.Println(instance)
		time.Sleep(time.Second)
	}
}
