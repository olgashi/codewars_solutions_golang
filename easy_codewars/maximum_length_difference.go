package main

import (
	"fmt"
	"math"
)

func findMaxLengthDifference(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	arrOneSmallest, arrOneLongest := findShortestAndLongestLength(a1)
	arrTwoSmallest, arrTwoLongest := findShortestAndLongestLength(a2)
	fmt.Println(arrOneSmallest, arrOneLongest)
	fmt.Println(arrTwoSmallest, arrTwoLongest)

	smallestDifference := math.Max(math.Abs(float64(arrOneSmallest)-float64(arrTwoLongest)), math.Abs(float64(arrOneLongest)-float64(arrTwoSmallest)))

	return int(smallestDifference)
}

func findShortestAndLongestLength(slice []string) (int, int) {
	shortest := slice[0]
	longest := slice[0]

	for _, element := range slice {
		if len(shortest) > len(element) {
			shortest = element
		}
		if len(longest) < len(element) {
			longest = element
		}
	}

	return len(shortest), len(longest)
}

func main() {
	fmt.Println(findMaxLengthDifference(
		[]string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"},
		[]string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}))
	fmt.Println(findMaxLengthDifference([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, []string{"bbbaaayddqbbrrrv"}))
}
