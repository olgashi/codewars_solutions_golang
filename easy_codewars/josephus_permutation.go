package main

import (
	"fmt"
)
/*
[1,_,_,4,5,_,7] - initial sequence
[1,2,4,5,6,7] => 3 is counted out and goes into the result [3]
[1,2,4,5,7] => 6 is counted out and goes into the result [3,6]
[1,4,5,7] => 2 is counted out and goes into the result [3,6,2]
[1,4,5] => 7 is counted out and goes into the result [3,6,2,7]
[1,4] => 5 is counted out and goes into the result [3,6,2,7,5]
[4] => 1 is counted out and goes into the result [3,6,2,7,5,1]
[] => 4 is counted out and goes into the result [3,6,2,7,5,1,4]

*/



func Josephus(items []interface{}, k int) []interface{} {
  itemsCopy := items[:]
	resultingSlice := []interface{}{}
	index := 0

	for len(itemsCopy) > 0 {
		if len(itemsCopy) == 1 {
			return append(resultingSlice, itemsCopy[0])
		}
		index = ((index + k) - 1) % len(itemsCopy)		
		resultingSlice = append(resultingSlice, itemsCopy[index])
		itemsCopy = append(itemsCopy[:index], itemsCopy[index+1:]...)
	}
	return resultingSlice
}


func main() {
	fmt.Println(Josephus([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1))
	fmt.Println(Josephus([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2))
	fmt.Println(Josephus([]interface{}{1, 2, 3, 4, 5, 6, 7}, 3))
	fmt.Println(Josephus([]interface{}{}, 3))
	
}