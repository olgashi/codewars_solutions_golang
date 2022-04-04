package main

import (
	"fmt"
)

func HighestRank(nums []int) int {
	var numCounts = make(map[int]int)
	for _, num := range nums {
		numCounts[num]++
	}
  maxNum := 0
	occurTimes := 0
	for k, v := range numCounts {
    if v > occurTimes || (v == occurTimes && k > maxNum ) {
			maxNum = k
			occurTimes = v
		}
	}
	return maxNum
}


func main() {
	fmt.Println(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12})) // 12
	fmt.Println(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12, 10})) // 12
	fmt.Println(HighestRank([]int{12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10})) // 3
	fmt.Println(HighestRank([]int{8, 8, 3, 3, 2, 4, 10, 12, 10, 1, 1, 1, 1})) // 1
}