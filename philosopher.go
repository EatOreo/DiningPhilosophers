package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Philosophize(n int, self, leftFork, rightFork Entity) {
	var eaten int
	var eating bool
	for true {
		query := Request{"take?", make(chan bool)}
		leftFork.Input <- query
		if <- query.Confirm {
			rightFork.Input <- query
			if <- query.Confirm {
				eating = true
				eaten++
				fmt.Println("Philosopher", n + 1, "started eating")
			} else {
				query.Msg = "putdown"
				leftFork.Input <- query
			}
		}
		if eating {
		 	time.Sleep(time.Duration(rand.Intn(int(time.Second) * 2)))
			query.Msg = "putdown"
			leftFork.Input <- query
			rightFork.Input <- query
		 	eating = false
			fmt.Println("Philosopher", n + 1, "finished eating")
		} 
		time.Sleep(time.Duration(int(time.Second) + rand.Intn(int(time.Second) * 2)))
	}
}
