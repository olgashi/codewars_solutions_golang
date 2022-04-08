package main

import (
	"fmt"
)

func FindSummands(n int) (output []int) {
	if n == 1 {
		return []int{1}
	}
	cube := n*n*n
	middle := cube / n
	start := middle - n + 1

	for num := start; len(output) < n; num +=2 {
		fmt.Println(num)
		output = append(output, num)
	}
	return
}

func main() {
  fmt.Println(FindSummands(3)) //[7, 9, 11]
  fmt.Println(FindSummands(4)) // [13, 15, 17, 19]
}