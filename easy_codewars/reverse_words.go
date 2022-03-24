package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) (output string) {
	strList := strings.Split(str, " ")
  for index, word := range strList {
		strList[index] = ReverseSingleWord(word)
	}
	return strings.Join(strList, " ")
}

func ReverseSingleWord(str string) (output string) {
	for _, letter := range str {
    output = string(letter) + output
	}
	return
}

func main() {
  fmt.Println(ReverseWords("This is an example!")) // "sihT si na !elpmaxe"
  fmt.Println(ReverseWords("double  spaced  words")) // "elbuod  decaps  sdrow"
}