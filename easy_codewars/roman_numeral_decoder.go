package main

import (
	"fmt"
	"strings"
)

func Decode(roman string) int {
	romanToDecimalMap := map[string]int{"I":1, "V":5, "X":10, "L":50, "C":100, "D":500, "M":1000}
	convertedNum := 0
	romanCharsList := strings.Split(roman, "")

	for index := 0; index < len(romanCharsList); index += 1 {
		if index == len(romanCharsList) - 1 {
			convertedNum += romanToDecimalMap[romanCharsList[index]]
		} else if romanToDecimalMap[romanCharsList[index]] < romanToDecimalMap[romanCharsList[index+1]]{
			convertedNum += romanToDecimalMap[romanCharsList[index]] * -1
		} else {
			convertedNum += romanToDecimalMap[romanCharsList[index]]
		}
	}
  return convertedNum
}

func main() {
	fmt.Println(Decode("XXI")) // 21
	fmt.Println(Decode("I")) // 1
	fmt.Println(Decode("IV")) // 4
	fmt.Println(Decode("XIII")) // 13
	fmt.Println(Decode("MCCVI")) // 1206
	fmt.Println(Decode("XLII")) // 42
}
