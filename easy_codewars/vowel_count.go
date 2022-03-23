package main

import (
	"fmt"
	"strings"
)

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