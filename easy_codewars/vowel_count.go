package main

import (
	"fmt"
	"strings"
)

// PEDAC
// vowels in our case are: a, e, i, o, u
// input string consists of lower case letters and spaces only

// GetCount("abcd") // 1
// GetCount("bcd") // 1
// GetCount("apple") // 2
// GetCount("yellow") // 2
// GetCount("yellow apple") // 4
// GetCount("aeio") // 4

// Data Structures:
// slice to contain a, e, i, o, u

// Algorithm:

// create slice of type string and assign {"a", "e", "i", "o", "u"} to it - vowelsList
// create a variable count and set to 0
// iterate over vowelsList, at each iteration
// count how many occurences of the current element (vowel) are there in the input string
// add that number to count

// return count

func GetCount(str string) (count int) {
	var vowelsList []string = []string{"a", "e", "i", "o", "u"}
	count = 0

	for _, vowel := range vowelsList {
		count += strings.Count(str, vowel)
	}

  return count
}

func main() {
	fmt.Println(GetCount("abcd")) // 1
	fmt.Println(GetCount("bcd")) // 0
	fmt.Println(GetCount("apple")) // 2
	fmt.Println(GetCount("yellow")) // 2
	fmt.Println(GetCount("yellow apple")) // 4
	fmt.Println(GetCount("aeio")) // 4 
}