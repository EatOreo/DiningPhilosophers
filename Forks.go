package main

import (
	"math/rand"
	"time"
	"strconv"
)

func Forkiphize(n int, self Entity) {
	var inUse bool
	var uses int
	for true {
		//The fact that there can only be one request in the input channel at a time guarantees that only one
		//philosopher can pick use a fork at a time
		req := <-self.Input
		switch req.Msg {
		case "take?":
			if !inUse {
				inUse = true
				//
				uses++
				//The request has to be confirmed or denied otherwise the philospher would be left in a deadlock
				//The request has its own channel for the fork to confirm the request instead of using its Output
				//channel which could cause a race condition, because another philosopher than the one who send
				//the request, could recieve the answer
				req.Confirm <- true
			} else {
				req.Confirm <- false
			}
		case "putdown":
			inUse = false
		case "state":
			if inUse {
				self.Output <- "Fork " + strconv.Itoa(n+1) + " is being used"
			} else {
				self.Output <- "Fork " + strconv.Itoa(n+1) + " is not being used"
			}
		case "times":
			self.Output <- "Fork " + strconv.Itoa(n+1) + " has been used " + strconv.Itoa(uses) + " times"
		case "all":
			if inUse {
				self.Output <- "in use, and has " + strconv.Itoa(uses) + " (attempted) uses"
			} else {
				self.Output <- "not in use, and has " + strconv.Itoa(uses) + " (attempted) uses"
			}
		default:
			self.Output <- "Not a query, try 'state' or 'times'"
		}
		//just a little delay (nanoseconds), not neccesarry
		time.Sleep(time.Duration(int(rand.Intn(10000000))))
	}
}
