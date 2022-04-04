package main

import (
	"fmt"
	"strings"
	"strconv"
)

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