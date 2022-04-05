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
	// fmt.Println(Pyramid(1)) // [[1]]
	// fmt.Println(Pyramid(2)) // [[1],[1,1]]
	// fmt.Println(Pyramid(3)) // [[1],[1,1],[1,1,1]]
	fmt.Println(Pyramid(4)) // [[1],[1,1],[1,1,1][1,1,1,1]]
}

/*
PEDAC
Input: integer n >=0, representing number of sub arrays and range of length betyween the first sub array and last subarray
Output: nested array of length n

- all sub arrays should be filled with 1s
- each subarray is filled with elements 
- if n is 0, no subarrays are added


DS: arrays (nested)
// fmt.Println(Pyramid(3)) // [[1],[1,1],[1,1,1]]

i= 3

Algorithm:
- create an array of size n, set to empty array - outputArray

- loop for i..1 to <= n, at each iteration:
  create an array of size i and fill it with i number of 1s
	append to output array

return outputArray
*/