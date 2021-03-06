package main

import (
	"fmt"
	// Don't forget to `go get golang.org/x/tour/tree`
	"golang.org/x/tour/tree"
)

func doTheWalk(tree *tree.Tree, channel chan int) {
	if tree == nil {
		return
	}
	doTheWalk(tree.Left, channel)
	channel <- tree.Value
	doTheWalk(tree.Right, channel)
}

func Walk(tree *tree.Tree, channel chan int) {
	doTheWalk(tree, channel)
	close(channel)
}

func Same(firstTree, secondTree *tree.Tree) bool {
	firstChannel, secondChannel := make(chan int), make(chan int)

	go Walk(firstTree, firstChannel)
	go Walk(secondTree, secondChannel)

	for {
		firstValue, firstOk := <-firstChannel
		secondValue, secondOk := <-secondChannel
		fmt.Println(firstValue)
		fmt.Println(secondValue)
		if !firstOk || !secondOk {
			return firstOk == secondOk
		}
		if firstValue != secondValue {
			return false
		}
	}
}

func main() {
	fmt.Print("tree.New(2) == tree.New(2): ")
	if Same(tree.New(2), tree.New(2)) {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
	fmt.Print("tree.New(1) != tree.New(2): ")
	if Same(tree.New(1), tree.New(2)) {
		fmt.Println("FALSE")
	} else {
		fmt.Println("TRUE")
	}
}
