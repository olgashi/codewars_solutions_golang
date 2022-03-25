package main

import (
	"fmt"
)

func Divisors(n int) (divisorCount int) {

	if n == 1 {
		return 1
	}

	endingDivisor := n

	for count := 1; count < endingDivisor; count++ {
		remainder := n % count
		if remainder == 0 {
			endingDivisor = n / count
			if endingDivisor != count {
				divisorCount += 2
			} else {
				divisorCount += 1
			}
		}
	}
	return divisorCount
}

func main() {
	fmt.Println(Divisors(4))  // 3 => 1, 2, 4
	fmt.Println(Divisors(5))  // 2
	fmt.Println(Divisors(12)) // 6
	fmt.Println(Divisors(30)) // 8
	fmt.Println(Divisors(1))  // 1
	fmt.Println(Divisors(10)) // 4
	fmt.Println(Divisors(11)) // 2
	fmt.Println(Divisors(54)) // 8
	fmt.Println(Divisors(64)) // 7

}
