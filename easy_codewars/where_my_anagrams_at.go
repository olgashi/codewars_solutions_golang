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




// create a copy of word, sort the copy - split, sort, join - inputWordCopy
// create an empty slice of type string - outputAnagramsSlice

// create a copy of words - wordsCopy
// for each word in words:
// 
// create a copy of word, sort the copy - split, sort, join
// if current word == as inputWordCopy
//   add to outputAnagramsSlice

// return outputAnagramsSlice
}


func main() {
	fmt.Println(Anagrams("abba", []string{"aabb", "abcd", "bbaa", "dada"}))
}