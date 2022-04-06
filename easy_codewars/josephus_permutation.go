package main

import (
	"fmt"
)

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