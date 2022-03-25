package main

import (
	"fmt"
)

func IsAscOrder(numbers []int) bool {
	for index := 1; index < len(numbers); index++ {
		if numbers[index] < numbers[index-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsAscOrder([]int{1, 2, 4, 7, 19}))            // true
	fmt.Println(IsAscOrder([]int{1, 2, 3, 4, 5}))             // true
	fmt.Println(IsAscOrder([]int{1, 6, 10, 18, 2, 4, 20}))    // false
	fmt.Println(IsAscOrder([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})) // false
	fmt.Println(IsAscOrder([]int{51}))                        // true
	fmt.Println(IsAscOrder([]int{89, 1}))                     // false
	fmt.Println(IsAscOrder([]int{8, 19}))                     // true
	fmt.Println(IsAscOrder([]int{8, 8}))                      // true
}
