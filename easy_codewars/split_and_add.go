package main

import (
	"fmt"
)

func SplitAndAdd(numbers []int, n int) []int {
	numbersCopy := numbers[:]
  addedTimes := 1
	for len(numbersCopy) > 1 && addedTimes <= n {
		result := []int{}
		middle := len(numbersCopy) / 2
		part1 :=  numbersCopy[:middle]
		part2 := numbersCopy[middle:] // always longer if length is odd
		
		if len(part1) < len(part2) {
			temp := part1[:]
			part1 = []int{0}
			part1 = append(part1, temp...)
		}

		for index := 0; index < len(part1); index++ {
			result = append(result, part1[index] + part2[index])
		}
    numbersCopy = result[:]
		addedTimes++
	}
	return numbersCopy
}

func main() {
	fmt.Println(SplitAndAdd([]int{1, 2, 3, 4, 5}, 2)) // [5, 10]
	fmt.Println(SplitAndAdd([]int{1, 2, 3, 4, 5}, 10)) // [15]
	fmt.Println(SplitAndAdd([]int{2, 2, 2, 2, 2, 2}, 10)) // [12]
	fmt.Println(SplitAndAdd([]int{2, 1, 2, 2, 1, 2}, 10)) // [10]
	fmt.Println(SplitAndAdd([]int{5, 1, 2, 5, 1, 2}, 1)) // [10, 2, 4]
}