package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// gives tun time exception as it is nil interface
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
