package main

import (
	"fmt"
)

func MakeNegative(x int) int {
	if x <= 0 {
		return x
	}
	return x * -1
}

func main() {
	fmt.Println(MakeNegative(-10)) // -10
	fmt.Println(MakeNegative(-7))  // -7
	fmt.Println(MakeNegative(0))   // 0
	fmt.Println(MakeNegative(51))  // -51
	fmt.Println(MakeNegative(78))  // -78
}
