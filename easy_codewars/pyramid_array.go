package main

import (
	"fmt"
)

func Pyramid(n int) [][]int {
 outputArray := [][]int{}

	for size := 0; size < n; size++ {
		subArr := []int{}
		for i := 0; i <= size; i++ {
			subArr = append(subArr, 1)
		}
		outputArray = append(outputArray, subArr)
	}
	return outputArray[:]
}

func main() {
	fmt.Println(Pyramid(0)) // []
	fmt.Println(Pyramid(1)) // [[1]]
	fmt.Println(Pyramid(2)) // [[1],[1,1]]
	fmt.Println(Pyramid(3)) // [[1],[1,1],[1,1,1]]
	fmt.Println(Pyramid(4)) // [[1],[1,1],[1,1,1][1,1,1,1]]
}
