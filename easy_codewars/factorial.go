package main

import (
	"fmt"
)

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n - 1)
}

func main() {
	fmt.Println(Factorial(4)) // 24 => 4 * 3 * 2 * 1
	fmt.Println(Factorial(1)) // 1
	fmt.Println(Factorial(7)) // 5040
	fmt.Println(Factorial(0)) // 1

}