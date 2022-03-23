package main

import (
	"fmt"
	"strings"
)

func RemoveChar(word string) string {
	if len(word) < 2 {
		return word
	}
	splitStr := strings.Split(word, "")
	endIndex := len(splitStr) - 1
	return strings.Join(splitStr[1:endIndex], "")
}

func main() {
	fmt.Println(RemoveChar("abc def")) // bcde
	fmt.Println(RemoveChar("bc")) // bc
	fmt.Println(RemoveChar("L")) // L
	fmt.Println(RemoveChar("UP")) // UP
	fmt.Println(RemoveChar("dowN and up")) // owN and u
	fmt.Println(RemoveChar("I")) // I
	fmt.Println(RemoveChar("a ")) // a
	fmt.Println(RemoveChar("")) // ""
}
