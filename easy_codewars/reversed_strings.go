package main
import (
	"fmt"
	"strings"
)

func Solution(word string) string {
	reversedWordSlice := make([]string, len(word))

	inputWordSlice := strings.Split(word, "")

	for index, letter := range inputWordSlice {
		reversedIndex := len(word) - (index + 1)
		reversedWordSlice[reversedIndex] = letter
	}
	
	return strings.Join(reversedWordSlice, "")
}

func main() {
	fmt.Println(Solution("")) // ""
	fmt.Println(Solution("a")) // a
	fmt.Println(Solution("abba")) // abba
	fmt.Println(Solution("world")) // dlrow
	fmt.Println(Solution("word")) // drow
	fmt.Println(Solution("cat")) // tac
}

// PEDAC

// Input: a string
// Output: an original string in a reversed form

// reverse the string provided as input
// Questions
// What if the input is an empty string? Or if the argument is missing? Or if there is more than one argument?

// Examples
//done

// Data Strucutres:
// slice

// {w, o, r, d} => 
// { d,r, o,w}
// Algorithm

// create a slice of length == len(input string)- reversedStr
// split the input string (effectively creating a slice from it) - inputStrSlice
// iterate over it, at each iteration
// place the current element into reversedStr at index == len(inputStrSlice) - (current index + 1)

// join the slice => string
// return the resulting string


