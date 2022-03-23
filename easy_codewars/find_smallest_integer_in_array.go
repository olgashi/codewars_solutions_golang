package main

import (
	"fmt"
	"math"
)

func SmallestIntegerFinder(numbers []int) int {
	smallest := math.Inf(0)
	for _, num := range numbers {
		if float64(num) < smallest {
			smallest = float64(num)
		}
	}
	return int(smallest)
}

func main() {
	fmt.Println(SmallestIntegerFinder([]int{34, 15, 88, 2})) // 2
	fmt.Println(SmallestIntegerFinder([]int{34, -345, -1, 100})) // -345
	fmt.Println(SmallestIntegerFinder([]int{34, 345, 0, 100})) // 0
}