package main

import (
	"fmt"
	"strings"
	"strconv"
)

func DigitalRoot(n int) int {
	reducedNumber := n
	var reducedNumberStr[]string = strings.Split(fmt.Sprintf("%d", reducedNumber), "")

	for len(reducedNumberStr) > 1 {
		var currentSum int
		for _, char := range reducedNumberStr {
			currentInt, _ := strconv.Atoi(char)
			currentSum += currentInt
		}
		reducedNumberStr = strings.Split(fmt.Sprintf("%d", currentSum), "")
	}

	finalNum, _ := strconv.Atoi(strings.Join(reducedNumberStr, ""))
	return finalNum
}


func main() {
	fmt.Println(DigitalRoot((16))) // 7
	fmt.Println(DigitalRoot((942))) // 6
	fmt.Println(DigitalRoot((132189))) // 6
	fmt.Println(DigitalRoot((493193))) // 2
	fmt.Println(DigitalRoot((581))) // 5
	fmt.Println(DigitalRoot((100079))) // 8
}
