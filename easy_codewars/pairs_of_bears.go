package main

import (
	"fmt"
)
func CheckPairs(nbBear int, bearList string) (bearCouple string, hasEnoughCouple bool) {

	ptr2 := 1
	for ptr1 := 0; ptr1 < len(bearList) - 1; ptr1++ {
		possiblePair := string(bearList[ptr1]) + string(bearList[ptr2])
		if possiblePair == "B8" || possiblePair == "8B" {
			bearCouple += possiblePair
			ptr1++
			ptr2 += 2
		} else {
			ptr2++
		}
	}

  return bearCouple, len(bearCouple) >= nbBear
}

func main() {
	fmt.Println(CheckPairs(7, "8j8mBliB8gimjB8B8jlB")) //"B8B8B8", false
	fmt.Println(CheckPairs(3, "88Bifk8hB8BB8BBBB888chl8BhBfd")) // "8BB8B8B88B", true
}