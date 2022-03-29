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
// PEDAC
/*
Problem:

Inputs: array of integers and a target number (integer)

Outputs: array of two integers, representing indices of numbers that when 
  added together sum up to the target number

- If there are more than one answer/solution, return the first one
- There always be at least two elements in array

Examples:


DS: none

Algorithm:
- create an empty map - uniqueNumIndexMap {1:0, 2:1, 4:3, 5:4, 9:5}

loop over input list, for each element, num1
  if el not in map
	  add it to map as a key and its index as a value
		calculate difference beteen target and number - num2
		if num2 in map, get its index (value ofr key num2) {
			return arr like: [num1 index, num2 index]
		}
Code
*/

func main() {
	fmt.Println(TwoSum([]int{1, 2, 3}, 4)) // [0, 2]
	fmt.Println(TwoSum([]int{1234, 5678, 9012}, 14690)) // [1, 2]
	fmt.Println(TwoSum([]int{2, 2, 3}, 4)) // [0, 1]
	fmt.Println(TwoSum([]int{1, 2, 2, 4, 5, 9, 7}, 14)) // [4, 5]
	fmt.Println(TwoSum([]int{0, 1, 2, 3, 0}, 0)) // 0 [0, 4]
}