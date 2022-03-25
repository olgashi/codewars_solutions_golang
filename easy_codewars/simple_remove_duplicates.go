package main

import (
	"fmt"
)

func removeDuplicates(arr []int) (uniqueIntsList []int) {
	uniqueIntsMap := make(map[int]bool)

	for index := len(arr) - 1; index >=0; index-- {
		num := arr[index]
		if _, numExists := uniqueIntsMap[num]; !numExists {
			uniqueIntsMap[num] = true
			uniqueIntsList = append([]int{num}, uniqueIntsList...)
		}
	}
	return uniqueIntsList
}

func main() {
	fmt.Println(removeDuplicates([]int{1,2,1,2,1,1,3}))
	fmt.Println(removeDuplicates([]int{0, 4, 4, 3, 0, 3}))
}