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
	balancer := balance.RandomBalance{}
	for {
		in, err := balancer.Dobalance(arr)
		if err == nil {
			fmt.Println(in)
		} else {
			fmt.Println(err.Error())
		}
		time.Sleep(time.Second)
	}
}
