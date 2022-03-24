package main

import (
	"fmt"
	"strings"
	"sort"
)


func TwoToOne(s1 string, s2 string) string {

	stringsSlice := strings.Split(s1 + s2, "")
	sort.Strings(stringsSlice)

	var uniqueLetters string

	for _, el := range stringsSlice {
		if !strings.Contains(uniqueLetters, el) {
			uniqueLetters += el
		}
	}

	return uniqueLetters
}

func main() {
	// Examples:
	fmt.Println(TwoToOne("xyaabbbccccdefww", "xxxxyyyyabklmopq")) // "abcdefklmopqwxy"
	fmt.Println(TwoToOne("abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz")) // "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(TwoToOne("aaaaa", "bbbbb")) // "ab"
	fmt.Println(TwoToOne("zzzzz", "bbbbblm")) // "blmz"
	fmt.Println(TwoToOne("abcdeef", "abbcdeeff")) // "abcdef"
	fmt.Println(TwoToOne("abbcdeeff", "zzx")) // "abcdefxz"
}