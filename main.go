package main

import (
	"fmt"
)

//Best query to get an overview is just "all", philosopher n can use forks n and n+1

type Entity struct {
	Input  chan Request
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
	for true {
		var what, query string
		var which int
		fmt.Scan(&what)
		if (what == "end") {
			break
		} else if (what == "all") {
			for i, v := range(philos) {
				v.Input <- Request{"all", nil}
				fmt.Println("phil:", i + 1, "is", <-v.Output)
			}
			fmt.Println("")
			for i, v := range(forks) {
				v.Input <- Request{"all", nil}
				fmt.Println("fork:", i + 1, "is", <-v.Output)
			}
			continue
		}
		fmt.Scan(&which, &query)
		which--
		switch what {
		case "p":
			philos[which].Input <- Request{query, nil}
			fmt.Println(<- philos[which].Output)
			break
		case "f":
			forks[which].Input <- Request{query, nil}
			fmt.Println(<- forks[which].Output)
			break
		}
	}
}
