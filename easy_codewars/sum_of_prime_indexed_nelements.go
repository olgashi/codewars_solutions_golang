package main

import (
	"fmt"
)

func Solve(arr []int) int {
	sum := 0
	for index := 0; index < len(arr); index++ {
		if isPrime(index) {
			sum += arr[index]
		}
	}
	return sum
}

func isPrime(num int) bool{
	if num < 2 {
		return false
	}

	for index := 2; index < num; index++ {
		if (num != index && num % index == 0) {
			return false
		}
	}
	return true
}

func main() {
  fmt.Println(Solve([]int{1,2,3,4,5,6})) //13
}
