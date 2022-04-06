package main

import (
	"fmt"
)
//0, 1, 2, 3, 0 => 0
func TwoSum(numbers []int, target int) [2]int {
	var indexMap = make(map[int]int)
	var complementMap = make(map[int]int)

	for index, num := range numbers {
		if num2, exists := complementMap[num]; exists {
			return [2]int{indexMap[num2], index}
		} else {
			indexMap[num] = index
			complementMap[target - num] = num
		}
	}
	return [2]int{-99, -99}
}

func main() {
	fmt.Println(TwoSum([]int{1, 2, 3}, 4)) // [0, 2]
	fmt.Println(TwoSum([]int{1234, 5678, 9012}, 14690)) // [1, 2]
	fmt.Println(TwoSum([]int{2, 2, 3}, 4)) // [0, 1]
	fmt.Println(TwoSum([]int{1, 2, 2, 4, 5, 9, 7}, 14)) // [4, 5]
	fmt.Println(TwoSum([]int{0, 1, 2, 3, 0}, 0)) // 0 [0, 4]
}