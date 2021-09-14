package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var philos [5]chan int 
//var forks [5]bool
var forks [5]Fork
var mu sync.Mutex
var output chan string

func main() {
	output = make(chan string)
	for i := 0; i < 5; i++ {
		forks[i] = Fork{false, make(chan int)}
		philos[i] = make(chan int)
		go forkiphize(forks[i].in)
		go Philosophize(i, philos[i])
	}
	for true {
		var qu0, qu1, qu2 int
		fmt.Scanln(&qu0, &qu1, &qu2)
		if (qu0 == -1) {
			break
		}
		if(qu2 == 0){
		  philos[qu0] <- qu1
		}else{
			forks[qu0].in <- qu1
		}


		fmt.Println(<- output)
	}
}

func TakeFork(f,p int) {
	forks[f].used = true
	forks[f].in <-p
}

func PutFork(f int) {
	forks[f].used = false
	forks[f].in <- -3
}

func CanTake(f int) bool {
	return !forks[f].used
}

func useforks(r int, n int){
	forks[n].in <- 42
	forks[r].in <- 42
}

func Philosophize(n int, input chan int) {
	var eaten int
	var eating bool
	r := (n + 1) % 5
	for true {
		mu.Lock()
		if CanTake(n) {
			TakeFork(n, n+10)
			if CanTake(r) {
				TakeFork(r,n+10)
				eating = true
				eaten++
				useforks(r,n)
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
