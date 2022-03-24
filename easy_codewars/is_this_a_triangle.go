package main

import (
	"fmt"
)

func IsTriangle(a, b, c int) bool {
  return a + b > c && b + c > a && a + c > b
}

func main() {
	fmt.Println(IsTriangle(5, 1, 2)) // false
	fmt.Println(IsTriangle(1, 2, 5)) // false
	fmt.Println(IsTriangle(5, 1, 5)) // true
	fmt.Println(IsTriangle(2, 2, 2)) // true
	fmt.Println(IsTriangle(-1, 2, 3)) // false
}

