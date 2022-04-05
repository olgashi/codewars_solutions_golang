package main

import (
	"fmt"
)


func ProductFib(prod uint64) (outputArr [3]uint64) {
	var fib1, fib2, prodOfTwoFibonacciNums uint64 = 0, 1, 0

	for prodOfTwoFibonacciNums < prod {
		prodOfTwoFibonacciNums = fib1 * fib2

		if prodOfTwoFibonacciNums == prod {
			outputArr = [3]uint64{fib1, fib2, 1}
		} else if prodOfTwoFibonacciNums > prod {
			outputArr = [3]uint64{fib1, fib2, 0}
	  }

		fib1, fib2 = fib2, fib1 + fib2
	}
	return
}

func main() {
	fmt.Println(ProductFib(714)) // {21, 34, 1} 21*34 = 714
	fmt.Println(ProductFib(800)) // {34, 55, 0} 34 * 55 = 1870
	fmt.Println(ProductFib(16)) // {5, 8, 0} 3*5=15 < 16 <5*8=40 
	fmt.Println(ProductFib(104)) // {8, 13, 1}
}
