package main

import (
	"fmt"
	"strings"
	"math"
)

func FirstNonRepeating(str string) string {
  letterCountMap := make(map[string]map[string]int)

  for index, el := range str {
		letter := strings.ToLower(string(el))
		if _, exists := letterCountMap[letter]; !exists {
			letterCountMap[letter] = make(map[string]int)
		}

		letterCountMap[letter]["index"] = index
		letterCountMap[letter]["count"]++
	}
  firstOccuringUniqueLetter := ""
	smallestIndex := math.Inf(1)

	for _, v := range letterCountMap {
		if smallestIndex > float64(v["index"]) && v["count"] == 1{
			smallestIndex = float64(v["index"])
			firstOccuringUniqueLetter = string(str[int(smallestIndex)])
		}
	}

	return firstOccuringUniqueLetter
}

func main() {
  fmt.Println(FirstNonRepeating("sTress")) // 'T'
	fmt.Println(FirstNonRepeating("stress")) // 't'
	fmt.Println(FirstNonRepeating("sttrerress")) // ""
	fmt.Println(FirstNonRepeating("abcd")) // "a"
	fmt.Println(FirstNonRepeating("aaabbccdde")) // "e"
}
