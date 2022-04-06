package main

import (
	"fmt"
	"strings"
)

func toWeirdCase(str string) (outputStr string) {

	for _, word := range strings.Split(str, " ") {
		for letterIndex, letter := range strings.Split(word, "") {
			if letterIndex == 0 || letterIndex % 2 == 0 {
				outputStr += strings.ToUpper(letter)
			} else {
				outputStr += strings.ToLower(letter)
			}
		}

	outputStr += " "

	}
	return strings.TrimRight(outputStr, " ")
}

func main() {
	fmt.Println(toWeirdCase("ABC")) // => returns "AbC"
	fmt.Println(toWeirdCase("String")) // => returns "StRiNg"
  fmt.Println(toWeirdCase("Weird string case")) // => returns "WeIrD StRiNg CaSe"
}