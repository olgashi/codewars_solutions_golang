package main

import (
	"fmt"
	// "math"
)

func FindNb(m int) int {
  intermediateValue := m
	for num := 1; intermediateValue > 0; num++ {
		intermediateValue -= num * num * num

		if intermediateValue == 0 {
			return num
		}
	}
	return -1
}

func main() {
	fmt.Println(FindNb(1071225)) // 45
	fmt.Println(FindNb(4183059834009)) // 2022
	fmt.Println(FindNb(91716553919377)) // -1
	fmt.Println(FindNb(24723578342962)) // -1

}