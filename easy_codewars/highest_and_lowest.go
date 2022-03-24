package main

import (
	"fmt"
	"strings"
	"strconv"
)

// PEDAC

// Problem
// Input: a string representing a list of space separated numbers
// Outputs: string representing highest and lowest numbers

// all numbers are confirmed to be valid int32
// there will be at least 1 number in string
// output must be 2 numbers with highest number first
// can have negative numbers

// Examples



// DS
// slice - convert input string into slice (list)

// Algorithm
// split input string using space as the delimiter => slice - inputNumList
// create variable min, set to pos Infinity -3
// create variable max, set to neg Infinity 5

// iterate over inputNumList, for each iteration:
// if int(current Number) < min
//  set min to currentNumber
// if int(current Number) > max
//  set max to current Number

// join max and min on space and return the resulting string

// Code

// fmt.Println(HighAndLow("1 2 -3 4 5")) // "5 -3"

func HighAndLow(in string) string {
	
	inputNumList := strings.Split(in, " ")
	min, _ := strconv.ParseInt(inputNumList[0], 10, 32)
	max, _ := strconv.ParseInt(inputNumList[0], 10, 32)

	for _, num := range inputNumList {
		currentNum, _ := strconv.ParseInt(num, 10, 32)

		if currentNum < min {
			min = currentNum
		}
		if currentNum > max {
			max = currentNum
		}
	}
  return strings.Join([]string{fmt.Sprint(max), fmt.Sprint(min)}, " ")
}


func main() {
  fmt.Println(HighAndLow("1 2 3 4 5")) // "5 1"
	fmt.Println(HighAndLow("1 2 -3 4 5")) // "5 -3"
	fmt.Println(HighAndLow("1 9 2 4 -5")) // "9 -5"
	fmt.Println(HighAndLow("1")) // "1 1"
	fmt.Println(HighAndLow("56")) // "56 56"
	fmt.Println(HighAndLow("0 -100")) // "0 -100"
	fmt.Println(HighAndLow("-9")) // "-9 -9"
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")) // "42 -9"
}