package main

import (
	"fmt"
	"strings"
)

func RemoveDuplicateWords(str string) string {
	uniqueWords := make(map[string]bool)
	var outputList []string

	for _, word := range strings.Split(str, " ") {
		if _, valExists := uniqueWords[word]; !valExists {
			uniqueWords[word] = true
			outputList = append(outputList, word)
		}
	}
	return strings.Trim(strings.Join(outputList, " "), " ")
}

func main() {
	fmt.Println(RemoveDuplicateWords("alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta")) //alpha beta gamma delta
	fmt.Println(RemoveDuplicateWords("my cat is my cat fat"))                                                            //"my cat is fat"
	fmt.Println(RemoveDuplicateWords("apples and oranges"))                                                              // "apples and oranges"
	fmt.Println(RemoveDuplicateWords("apples and oranges apples and oranges"))                                           // "apples and oranges"
	fmt.Println(RemoveDuplicateWords("aaaapppplllles"))                                                                  // "aaaapppplllles"
}
