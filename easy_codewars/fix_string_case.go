package main

import (
	"fmt"
	"strings"
)

func fixStringCase(str string) string {
	lower, upper := 0, 0

	for _, letter := range str {
		if string(letter) == strings.ToLower(string(letter)) {
			lower += 1
		} else {
			upper += 1
		}
	}

	if upper == 0 || lower == 0 {
		return str
	} else if upper == lower || upper < lower {
		return strings.ToLower(str)
	}

	return strings.ToUpper(str)
}

func main() {
	fmt.Println(fixStringCase("a"))     // a
	fmt.Println(fixStringCase("B"))     // B
	fmt.Println(fixStringCase("coDe"))  // code
	fmt.Println(fixStringCase("CODe"))  // CODE
	fmt.Println(fixStringCase("coDE"))  // code
	fmt.Println(fixStringCase("apple")) // apple
	fmt.Println(fixStringCase("ZOO"))   // ZOO
}
