package mai

import (
	"fmt"
	"strings"
)


func ValidParentheses(parens string) bool {
	parensCount := 0

	for index := 0; index < len(parens); index++ {
    if parensCount < 0 {
			return false
		}
		if string(parens[index]) == "(" {
			parensCount++
		} else if string(parens[index]) == ")" {
			parensCount--
		} else {
			return false
		}
	}
	return parensCount == 0
}

func main() {

}

/*
PEDAC

Problem
-------

Inputs: 
string of parentheses (length of between 0 and 100) - "(" or ")"

Output:
boolean indicating if string of parentheses is valid

- valid string of parentheses - a string where each paren has a closing companion, and in correct order
for example:, "()()", "(())()()(())"
- invalid: "()(())(", ")("

Examples
-------

fmt.Println("()") // true
fmt.Println(")(()))") // false
fmt.Println("(") // false 
fmt.Println("(())((()())())") // true 
fmt.Println("()()(())") // true
fmt.Println("") // false
fmt.Println(")))") // false
fmt.Println("(()))(") // false

count

DS
-------
none

Algorithm
-------
// ( ) ( ) ( ( ) )

create a variable parensCount, set to 0

loop for index from 0 til the end of the string 
  if parensCount < 0 
	  return false

return parensCount == 0

Code
-------



*/