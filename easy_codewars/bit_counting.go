package main

import (
	"fmt"
)

func countBits(num uint) int {
	var bitCount int

	currentResult := num
	for currentResult > 0 {
		remainder := currentResult % 2
		currentResult = currentResult / 2

		if fmt.Sprintf("%d", remainder) == "1" {
			bitCount++
		}
	}

	return bitCount
}

func main() {
	fmt.Println(countBits(7))
}
