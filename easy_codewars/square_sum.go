package main

import (
	"fmt"
)

func SquareSum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num * num
	}
	return sum
}

func main() {
	fmt.Println(SquareSum([]int{1, 2, 2})) // 9
	fmt.Println(SquareSum([]int{1, 2, 3})) // 14
	fmt.Println(SquareSum([]int{2, 2, 2, 2})) // 16
}