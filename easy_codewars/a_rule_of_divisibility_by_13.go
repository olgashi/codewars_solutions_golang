package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Thirt(n int) int {
	isStationary := false
	sequenceIndex := 0
	previousSum := n
	sequence := []int{1, 10, 9, 12, 3, 4}

	for !isStationary {
		numList := strings.Split(strconv.Itoa(previousSum), "")
		currentSum := 0

		for index := len(numList) - 1; index >= 0; index-- {
			convInt, _ := strconv.Atoi(numList[index])
			currentSum += convInt * sequence[sequenceIndex]
			sequenceIndex++

			if sequenceIndex > len(sequence)-1 {
				sequenceIndex = 0
			}
		}

		if currentSum == previousSum {
			isStationary = true
			break
		} else {
			previousSum = currentSum
			sequenceIndex, currentSum = 0, 0
		}
	}
	return previousSum
}

func main() {
	fmt.Println(Thirt(1234567)) // 87
	fmt.Println(Thirt(321))     // 48
	fmt.Println(Thirt(512))     // 57
	fmt.Println(Thirt(102))     //11
	fmt.Println(Thirt(38))      // 38
}
