package main

import (
	"fmt"
)

func EvenOrOdd(number int) string {
	if number % 2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func main() {
  fmt.Println(EvenOrOdd(6))
  fmt.Println(EvenOrOdd(0))
  fmt.Println(EvenOrOdd(1))
  fmt.Println(EvenOrOdd(-10))
}