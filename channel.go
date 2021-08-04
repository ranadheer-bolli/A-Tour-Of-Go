package main

import (
	"fmt"
)

func sum(a []int, c chan int) {
	res := 0
	for _, v := range a {
		res += v
	}
	c <- res
}

func main() {

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}
