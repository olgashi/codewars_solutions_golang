package main

import (
	"fmt"
	"sort"
)


func Comp(array1[]int, array2 []int) bool {
  if array1 == nil || array2 == nil || len(array1) != len(array2) {
    return false
  }

	for idx, el := range array1 {
		array1[idx] = el*el
	} 

	sort.Ints(array1)
	sort.Ints(array2)

	for index, num := range array1 {
		if array2[index] != num {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Comp([]int{121, 144, 19, 161, 19, 144, 19, 11}, []int{11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19})) // true
	fmt.Println(Comp([]int{121, 144, 19, 161, 19, 144, 19, 22}, []int{11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19})) // false
	fmt.Println(Comp([]int{121, 19, 161, 19, 144, 19, 11}, []int{11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19})) // false
	fmt.Println(Comp([]int{}, []int{4, 6, 8})) // false
	fmt.Println(Comp([]int{2, 4}, []int{4, 16})) // true
	fmt.Println(Comp([]int{5, 4}, []int{16, 25})) // true
}

