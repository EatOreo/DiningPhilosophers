package main
import (
	"strconv"
)
type Fork struct{
	used bool
	in chan int
}

func forkiphize(inp chan int){
	uses := 0
	usedBy :=-1

	for true{
		select{
		case a:= <- inp:
			if(a == 42){
				uses++
			}else if(a == 0){
				output <- strconv.Itoa(uses)
			}else if(a >=10 && a<15){
				usedBy = a-10
			}else if(a == -3){
				usedBy = -1;
			}else if(a == 1){
				if(usedBy == -1){
					output <-"Not in use"
				}else{
					output <-strconv.Itoa(usedBy)
				}
			}
		default:
		}
}
}