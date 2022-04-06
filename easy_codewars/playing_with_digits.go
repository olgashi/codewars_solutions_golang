package main

import (
	"fmt"
	"strings"
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	var sum float64

	for index, element := range strings.Split(strconv.Itoa(n), "") {
		numberAsFloat, _ := strconv.ParseFloat(element, 64)
		sum += math.Pow(numberAsFloat, float64(index + p))
	}

	if int(sum) % n == 0 {
		result := int(sum) / n
		if result > 0 {
			return result
		}
	}
	return -1
}

func main() {
	fmt.Println(DigPow(89,1)) // 1
	fmt.Println(DigPow(92, 1)) // -1
	fmt.Println(DigPow(695, 2)) // 2
	fmt.Println(DigPow(46288, 3)) // 51
}