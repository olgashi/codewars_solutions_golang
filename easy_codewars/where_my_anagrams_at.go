package main

import (
	"fmt"
	"sort"
	"strings"
)

func Anagrams(word string, words []string) string {
	wordCopy := word

  sort.Strings(strings.Split(wordCopy, ""))
	return wordCopy
}


func main() {
	fmt.Println(Anagrams("abba", []string{"aabb", "abcd", "bbaa", "dada"}))
}