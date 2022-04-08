package main

import (
	"fmt"
	"math"
)

func Solve(n int) int {
	for i := 1; i < n; i++ {
		possiblePerfectSquare := (i * i) + n
		sqRoot := math.Sqrt(float64(possiblePerfectSquare))

		if sqRoot * sqRoot == float64(possiblePerfectSquare) && math.Trunc(sqRoot) == sqRoot {
		  return i * i
		}
	}
	return -1
}

func main() {
	fmt.Println(Solve(1)) // -1
	fmt.Println(Solve(2)) // -1
	fmt.Println(Solve(3)) // 1
	fmt.Println(Solve(4)) // -1
	fmt.Println(Solve(5)) // 4
	fmt.Println(Solve(7)) // 9
	fmt.Println(Solve(8)) // 1
	fmt.Println(Solve(9)) // 16
	fmt.Println(Solve(10)) // -1
	fmt.Println(Solve(11)) // 25
	fmt.Println(Solve(13)) // 36
	fmt.Println(Solve(17)) // 64
	fmt.Println(Solve(88901)) // 5428900
	fmt.Println(Solve(290101)) // 429235524
}