package main

import (
	"fmt"
	"strings"
)

func NbDig(n int, d int) (digitCount int) {

	for num := 0; num <= n; num++ {
		digitCount += strings.Count(fmt.Sprint(num * num), fmt.Sprint(d))
	}
	return
}

func main() {
	fmt.Println(NbDig(10, 1)) // 4 => the k*k are 0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100
	fmt.Println(NbDig(25, 1)) // 11 => 1, 16, 81, 100, 121, 144, 169, 196, 361, 441
	fmt.Println(NbDig(4, 8)) // 0 => 1, 4, 9, 16
	fmt.Println(NbDig(7, 9)) // 2 => 1, 4, 9, 12, 25, 36, 49
}