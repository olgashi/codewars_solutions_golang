package main

import (
	"fmt"
)

func century(year int) int {
	
	if remainder := year % 100; remainder > 0 {
		return (year / 100) + 1
	}  

	return year / 100
}

func main() {
	fmt.Println(century(1705)) // 18
  fmt.Println(century(1900)) // 19
  fmt.Println(century(1601)) //17
  fmt.Println(century(2000)) // 20
  fmt.Println(century(200000)) // 2000
  fmt.Println(century(0)) // 0
}