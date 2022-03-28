package main

import (
	"fmt"
)

func FindOdd(seq []int) int {
	var numCounts = make(map[int]int)
	
	for _, num := range seq {
		numCounts[num]++
	}

	for k, v := range numCounts {
		if (v % 2 != 0) {
			return k
		}
	}
	return 0
}


func main() {
	fmt.Println(FindOdd([]int{7}))
	fmt.Println(FindOdd([]int{0}))
	fmt.Println(FindOdd([]int{1, 1, 2}))
	fmt.Println(FindOdd([]int{0, 1, 0, 1, 0}))
	fmt.Println(FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))
	
}