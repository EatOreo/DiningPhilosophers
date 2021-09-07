package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var philos [5]chan int 
var forks [5]bool
var mu sync.Mutex
var output chan string

func main() {
	output = make(chan string)
	for i := 0; i < 5; i++ {
		philos[i] = make(chan int)
		go Philosophize(i, philos[i])
	}
	for true {
		var qu0, qu1 int
		fmt.Scanln(&qu0, &qu1)
		if (qu0 == -1) {
			break
		}

		philos[qu0] <- qu1

		fmt.Println(<- output)
	}
}

func TakeFork(f int) {
	forks[f] = true
}

func PutFork(f int) {
	forks[f] = false
}

func CanTake(f int) bool {
	return !forks[f]
}

func Philosophize(n int, input chan int) {
	var eaten int
	var eating bool
	r := (n + 1) % 5
	for true {
		mu.Lock()
		if CanTake(n) {
			TakeFork(n)
			if CanTake(r) {
				TakeFork(r)
				eating = true
				eaten++
			} else {
				PutFork(n)
			}
		} else {
		}
		mu.Unlock()
		select {
		case i := <- input:
			if (i == 0) {
				output <- strconv.Itoa(eaten)
			} else if (i == 1) {
				if (eating) {
					output <- "eating"
				} else {
					output <- "thinking"
				}
			} else {
				output <- "not a query"
			}
		default:
		}
		if eating {
			time.Sleep(time.Duration(rand.Intn(1000000000)))
			PutFork(n)
			PutFork(r)
			eating = false
		} 
		time.Sleep(time.Duration(rand.Intn(50000)))
	}
}
