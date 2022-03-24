package main

import "fmt"

func BreakChocholate(n, m int) int {
	if n < 1 || m < 1 {
		return 0
	}

	return (n * m) - 1
}

func main() {
	fmt.Println(BreakChocholate(5, 5)) // 24
	fmt.Println(BreakChocholate(2, 1)) // 1
	fmt.Println(BreakChocholate(2, 2)) // 3

}
