package main

import (
	"fmt"
)

func Tribonacci(signature [3]float64, n int) []float64 {
	var signatureCopy []float64
	signatureCopy = signature[:]
	if len(signature) < 3 {
		return signatureCopy
	}
	if n == 0 {
		return []float64{}
	}
	if n < 3 {
		return signatureCopy[:n]
	}

	for count := 0; count < n; count++ {
		signatureCopy = append(signatureCopy, signatureCopy[count] + signatureCopy[count + 1] + signatureCopy[count + 2])
	}
	return signatureCopy
}

func main() {
	fmt.Println(Tribonacci([3]float64{1, 1, 1}, 10)) //1, 1, 1, 3, 5, 9, 17, 31, 57, 105
	fmt.Println(Tribonacci([3]float64{0, 0, 1}, 10)) //0, 0, 1, 1, 2, 4, 7, 13, 24, 44
	fmt.Println(Tribonacci([3]float64{0, 1, 1}, 10)) //0, 1, 1, 2, 4, 7, 13, 24, 44, 81
	fmt.Println(Tribonacci([3]float64{-2, -1, 0}, 5)) // -3, -4, , -7, -14
	fmt.Println(Tribonacci([3]float64{0, 0, 0}, 3)) // 0, 0, 0
	fmt.Println(Tribonacci([3]float64{1, 2, 3}, 4)) // 6, 11, 20, 27
	fmt.Println(Tribonacci([3]float64{1, 2}, 1)) // 1
}
