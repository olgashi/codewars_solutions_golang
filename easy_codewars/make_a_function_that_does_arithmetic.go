package main

import (
	"fmt"
)

func Arithmetic(a int, b int, operator string) int {

	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		return a / b
	}
	return 0
}

func main() {
	fmt.Println(Arithmetic(5, 2, "add")) // 7 
	fmt.Println(Arithmetic(5, 2, "subtract")) // 3
	fmt.Println(Arithmetic(5, 2, "multiply")) // 10
	fmt.Println(Arithmetic(5, 2, "divide")) // 2

}