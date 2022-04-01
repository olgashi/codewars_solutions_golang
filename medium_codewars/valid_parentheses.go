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
	fmt.Println("()") // true
	fmt.Println(")(()))") // false
	fmt.Println("(") // false 
	fmt.Println("(())((()())())") // true 
	fmt.Println("()()(())") // true
	fmt.Println("") // false
	fmt.Println(")))") // false
	fmt.Println("(()))(") // false
}
