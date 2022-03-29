package main

import (
	"fmt"
	"strings"
)

func duplicate_count(s1 string) (letterCount int) {

	var charCountMap = make(map[string]int)
	for _, char := range strings.ToLower(s1) {
		letter := string(char)

		if charCountMap[letter]++; (charCountMap[letter] == 2) {
			letterCount++
		}
	}
	return letterCount
}

func main() {
	fmt.Println(duplicate_count("abcde")) // 0
	fmt.Println(duplicate_count("aabbcde")) // 2
	fmt.Println(duplicate_count("aabBcde")) // 2
	fmt.Println(duplicate_count("indivisibility")) // 1
	fmt.Println(duplicate_count("Indivisibilities")) // 2
	fmt.Println(duplicate_count("aA11")) // 2
	fmt.Println(duplicate_count("ABBA")) // 2
	fmt.Println(duplicate_count("AbbA")) // 2
	fmt.Println(duplicate_count("1100997744")) // 5
}