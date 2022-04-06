package main

import (
	"fmt"
	"math"
)

func Litres(time float64) int {
	const LITRES_PER_HOUR = 0.5

	return int(math.Floor(time * LITRES_PER_HOUR))
}

func  main() {
	fmt.Println(Litres(3)) // 1 => 3 * 0.5 => 1.5 => 1
	fmt.Println(Litres(6.7)) // 3 => 6.7 * 0.5 => 3.35 => 3
	fmt.Println(Litres(11.8)) // 5 => 11.8 * 0.5 => 5.9 => 5
	fmt.Println(Litres(0)) // 0
	fmt.Println(Litres(427)) // 213 => 427 * 0.5 => 213.5 => 213
	fmt.Println(Litres(6.78)) // 3 => 6.78 * 0.5 => 3.39 => 3
}