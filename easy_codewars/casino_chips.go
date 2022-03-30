package main

import (
	"fmt"
	"sort"
)

func Solve(arr []int) int {
	pilesOfChips := arr[:]
	daysCount := 0
	
	for len(pilesOfChips) > 1 {
		sort.Ints(pilesOfChips)

		if arr[len(arr) - 1] > 0 && arr[len(arr) - 2] > 0 {
			arr[len(arr) - 1]--
			arr[len(arr) - 2]--
			daysCount++
		}
		for len(pilesOfChips) > 0 && pilesOfChips[0] == 0 {
			pilesOfChips = pilesOfChips[1:]
		}	
	}
	return daysCount
}

/* Algorithm:

repeat until 1 pile left (two of the piles have value of 0)
  pick 2 piles that have largest (or equal to third) number of chips
  subtract 1 from each, 
  increment number of days

return number of days
*/
func main() {
  fmt.Println(Solve([]int{1, 1, 1}))	
  fmt.Println(Solve([]int{1, 2, 1}))	
  fmt.Println(Solve([]int{4, 1, 1}))	
  fmt.Println(Solve([]int{8, 2, 8}))	
}