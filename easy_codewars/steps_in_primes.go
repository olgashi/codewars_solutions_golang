package main

import (
	"fmt"
	"math"
)

func Step(step, start, end int) (result []int) {

	for num := start; num <= end - step; num++ {
		if isPrime(num) && isPrime(num + step) {
			return []int{num, num + step}
		}
	}
  return
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
  for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Step(2, 5, 7)) // [5, 7]
	fmt.Println(Step(2, 5, 5)) // nil
	fmt.Println(Step(4, 130, 200)) // [5, 7]

}