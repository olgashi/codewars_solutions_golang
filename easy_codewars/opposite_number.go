package main

import (
	"fmt"
)

func Opposite(x int) int {
	return x * -1
}

func main() {
	fmt.Println(Opposite(-10)) // 10
	fmt.Println(Opposite(-7))  // 7
	fmt.Println(Opposite(0))   // 0
	fmt.Println(Opposite(51))  // -51
	fmt.Println(Opposite(78))  // -78
}
