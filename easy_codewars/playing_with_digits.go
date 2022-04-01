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

/*
Problem 

Inputs:
- positive integer n 
- positive integer p ((starting power)

Output:
- integer k, such as: a ^ p + b ^ (p+1) + c ^(p+2) + d ^ (p+3) + ...) = n * k

Rules:
- find a positive integer k, if exists, such that the sum of the digits of n, 
taken to the successive powers of p is equal to k*n

- if doesnt exist return -1

Examples


DS
- slice - to be able to work with digits individually

Algorithm
//	fmt.Println(DigPow(92, 1)) // -1

- convert n to string, split on on empty space as the delimiter - save resulting slice to var inputNumsList

-initialize variable sum, set to 0
- loop over inputNumsList, for each element:
 - take it to the power of current elements index+p, add to sum

take sum and divide by input number n
if result is a whole number 
  return that number
else 
  return -1


Code



*/