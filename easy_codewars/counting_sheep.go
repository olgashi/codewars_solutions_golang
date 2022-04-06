package main

import (
	"fmt"
)

func CountSheeps(numbers []bool) (presentSeep int) {
	for _, el := range numbers {
		if el {
			presentSeep++
		}
	}
  return presentSeep
}

func main() {
	fmt.Println(CountSheeps([]bool{
		true,  true,  true,  false, 
		true,  true,  true,  true, 
		true,  false, true,  false,
		true,  false, false, true,
		true,  true,  true,  true,
		false, false, true,  true}))
}