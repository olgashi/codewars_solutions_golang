package main

import (
	"fmt"
)

func IsDivisible(num, x, y int) bool {
	return num % x == 0 && num % y == 0
}

func main() {
	fmt.Println(IsDivisible(3, 1, 3)) // true
	fmt.Println(IsDivisible(12, 2, 6)) // true
	fmt.Println(IsDivisible(100, 5, 3)) // false
	fmt.Println(IsDivisible(12, 7, 5)) // false
}