package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1009 * time.Millisecond)
	boom := time.After(5009 * time.Millisecond)
	for {
		select {
		case i := <-tick:
			fmt.Println("tick.", i)
		case i := <-boom:
			fmt.Println("BOOM!", i)
			return
		}

	}
}
