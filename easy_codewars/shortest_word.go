package main

import (
	"fmt"
	"strings"
)

func FindShort(s string) int {
	smallestWordLen := len(s)

	for _, word := range strings.Split(s, " ") {
		if len(word) < smallestWordLen {
			smallestWordLen = len(word)
		}
	}

	return smallestWordLen
}

func main() {
	fmt.Println(FindShort("a")) // 1
	fmt.Println(FindShort("ab")) // 2
	fmt.Println(FindShort("a ab abc a abcd a")) // 1
	fmt.Println(FindShort("aaba abcd a")) // 1
	fmt.Println(FindShort("at at at at att")) // 2
	fmt.Println(FindShort("word cat limo")) // 3
}