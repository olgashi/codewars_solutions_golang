package main

import (
	"fmt"
	"math"
)

// PEDAC

// Problem
// Input: float64 representing time of cycling
// Output: int representing liters of water

// 0.5 l == 1 hr , 0.5liters/hr

// return value of number of liters should be rounded to the smallest whole value (rounded down)

// Examples
// fmt.Println(Liters(3)) // 1 => 3 * 0.5 => 1.5 => 1
// fmt.Println(Liters(6.7)) // 3 => 6.7 * 0.5 => 3.35 => 3
// fmt.Println(Liters(11.8)) // 5 => 11.8 * 0.5 => 5.9 => 5
// fmt.Println(Liters(0)) // 0
// fmt.Println(Liters(427)) // 213 => 427 * 0.5 => 213.5 => 213
// fmt.Println(Liters(6.78)) // 3 => 6.78 * 0.5 => 3.39 => 3

// DS
// N/a

// Algorithm
// multiply imput int by 0.5 -> totalLiters
// round totalLiters down to get whole number of litres

// Code

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