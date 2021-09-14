package main

import (
	//"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Philosophize(n int, self, leftFork, rightFork Entity) {
	var eaten int
	var eating bool
	for true {
		req := Request{"take?", make(chan bool)}
		leftFork.Input <- req
		if <-req.Confirm {
			rightFork.Input <- req
			if <-req.Confirm {
				eating = true
				eaten++
				//fmt.Println("Philosopher", n+1, "started eating")
			} else {
				req.Msg = "putdown"
				leftFork.Input <- req
			}
		}
		do := true
		min := 1
		max := 2
		if eating {
			max = 4
		}
		finished := make(chan bool)
		go timeOut(min, max,finished)

		for do {
			select {
			case query := <-self.Input:
				switch query.Msg {
				case "state":
					if eating {
						self.Output <- "Philosopher " + strconv.Itoa(n + 1) + " is eating"
					} else {
						self.Output <- "Philosopher " + strconv.Itoa(n + 1) + " is thinking"
					}
				case "times":
					self.Output <- "Philosopher " + strconv.Itoa(n + 1) + " has eaten " + strconv.Itoa(eaten) + " times"
				case "all":
					if eating {
						self.Output <- "eating, and has " + strconv.Itoa(eaten) + " times"
					} else {
						self.Output <- "not eating, and has " + strconv.Itoa(eaten) + " times"
					}
				default:
					self.Output <- "Not a query, try 'state', 'times' or 'all'"
				}
			case fin := <-finished:
				do = !fin
			default:
			}
		}

		if eating {
			req.Msg = "putdown"
			leftFork.Input <- req
			rightFork.Input <- req
			eating = false
			//fmt.Println("Philosopher", n+1, "finished eating")
		}
		//just so a philosopher can't reclaim the forks instantly
		time.Sleep(time.Duration(int(rand.Intn(50000000))))
	}
}

func timeOut(min, max int, finished chan bool) {
	time.Sleep(time.Duration(min + rand.Intn(max - min) * int(time.Second)))
	finished <- true
}