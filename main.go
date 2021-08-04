package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
