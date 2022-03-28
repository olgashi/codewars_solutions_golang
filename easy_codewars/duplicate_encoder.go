package main

import (
	"fmt"
	"strings"
)

func DuplicateEncode(word string) string {

	charCounts := make(map[string]string)
	wordCharsList := strings.Split(word, "")
	var outputStr string

	for _, char := range wordCharsList {
		lowerChar := strings.ToLower(char)
		if _, ok := charCounts[lowerChar]; ok {
			charCounts[lowerChar] = ")"
		} else {
			charCounts[lowerChar] = "("
		}
	}

	for _, char := range wordCharsList {
		lowerChar := strings.ToLower(char)
		outputStr += charCounts[lowerChar]
	}

	return outputStr
}


func main() {
	fmt.Println(DuplicateEncode("din"))
	fmt.Println(DuplicateEncode("recede"))
	fmt.Println(DuplicateEncode("Success"))
	fmt.Println(DuplicateEncode("(( @"))
	
}