package main

import (
	"fmt"
	"strings"
)

func ValidBraces(str string) bool {
	bracesMap := map[string]string{"(":")","[":"]", "{":"}"}
	inputStrSlice := strings.Split(str, "")

	for {
    index := 0
		fullSet := false
		for index + 1 < len(inputStrSlice) {
			if bracesMap[inputStrSlice[index]] == inputStrSlice[index+1] {
				inputStrSlice = append(inputStrSlice[:index], inputStrSlice[index+2:]...)
				index = 0
				fullSet = true
			} else {
				index++
			}
		}
		if len(inputStrSlice) == 0 {
			return true
		}
		if !fullSet {
			return false
		}
	}
}

func main() {
	fmt.Println(ValidBraces("(){}[]")) // true
	fmt.Println(ValidBraces("([{}])")) //true
	fmt.Println(ValidBraces("(}")) // false
	fmt.Println(ValidBraces("[(])")) // false
	fmt.Println(ValidBraces("[({})](]")) // false
	fmt.Println(ValidBraces("[{}](){}")) // true
	fmt.Println(ValidBraces("[[{}]]()")) // true
	fmt.Println(ValidBraces("((}")) // false	
}

/*
Problem
-------

Inputs:
a string, containing (), [], {}
Outputs:
boolean t/f representing if all braces are valid or invalid

Objective: determine if the input string is valid

- valid string - when all braces are matched with the correct brace: (){}[], ([{[]}])
- input string will only contsain: (), {}, []

Examples:



DS:
slice (result of splitting the input string)
map - k - opening brace, v matching closing brace

// case: "(}"

Algorithm:
- map contaning {}, [], () as k/v - bracesMap
- split the input string - inputStrSlice
index = 0

loop over inputStrSlice (while true):

  loop while index + 1 < len(inputStrSlice) {
		if bracesMap[inputStrSlice[index]] == inputStrSlice[index + 1] {
			delete pair from inputStrSlice
			set index = 0
			break
		} else {
			index++
		}
	}

	if len(inputStrSlice) == 0 
	  return true


	return false
	
Code

*/