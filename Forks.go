package main

import (
)

func Forkiphize(n int, self Entity){
	var inUse bool
	var uses int

	for true {
		query := <- self.Input
		//fmt.Println(n, query)
		switch query.Msg {
		case "take?":
			if !inUse {
				inUse = true
				uses++
				query.Confirm <- true
				} else {
					query.Confirm <- false
				}
		case "putdown":
			inUse = false
		}		
	}
}
