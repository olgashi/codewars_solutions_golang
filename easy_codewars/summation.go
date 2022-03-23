package main

import (
	"fmt"
)

func Summation(n int) int {
	if n == 1 {
		return n
	}
	return n + Summation(n-1)
}

func main() {
	fmt.Println(Summation(2)) // 3
	fmt.Println(Summation(5)) // 15
}