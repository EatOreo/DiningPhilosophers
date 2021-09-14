package main

import "time"

type Entity struct {
	Input chan Request
	Output chan string
}

type Request struct {
	Msg     string
	Confirm chan bool
}

func main() {
	var philos, forks [5]Entity
	for i := 0; i < 5; i++ {
		philos[i].Input = make(chan Request)
		philos[i].Output = make(chan string)
		forks[i].Input = make(chan Request)
		forks[i].Output = make(chan string)
	}
	for i := 0; i < 5; i++ {
		go Philosophize(i, philos[i], forks[i], forks[(i+1)%5])
		go Forkiphize(i, forks[i])
	}
	time.Sleep(time.Second * 10)
}
