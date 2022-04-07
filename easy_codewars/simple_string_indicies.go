package main

import (
	"fmt"
	"errors"
)

func Solution(str string, idx uint) (uint, error) {
	openingParenIndxList := []int{}

	for index, el := range str {
		if string(el) == "(" {
			openingParenIndxList = append(openingParenIndxList, index)
		} else if string(el) == ")" {
			if openingParenIndxList[len(openingParenIndxList) - 1] == int(idx) {
				return uint(index), nil
			} else {
				openingParenIndxList = openingParenIndxList[:len(openingParenIndxList)-1]
			}
		}
	}
	
  return 0, errors.New("Not an opening bracket")
}

func main() {
  fmt.Println(Solution("((1)23(45))(aB)", 0))
}
