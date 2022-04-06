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