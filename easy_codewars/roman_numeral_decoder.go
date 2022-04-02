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


/*

PEDAC


Problem
- Input: a string representing a roman numeral
- Output: an integer representing a decoded roman numeral

Roman numerals are written by expressing each digit encoded separately 

- if smaller number (roman) comes before a larger number, 
smaller needs to be subtracted from larger, and then added to running total

- otherwise each number added to a running total


Examples: done

XXI // 21 as 
X = 10
X = 10

IV = 4

MMVIII = 2008

MDC L X V I => 1000 + 500 +100 + 50 + 10 + 5 + 1 = 1666

Symbol    Value
I          1
V          5
X          10
L          50
C          100
D          500
M          1,000

DS:
map where key is a roman numeral and value is its decimal equivalent
slice - to hold all input string split into characters

Algoritm:
 - create variable to hold the output - convertedNum
- create map where key is a roman numeral and value is its decimal equivalent - romanToDecimalMap
- split input string on empty space, save in variable romanCharsList
- loop over each charachter starting with the second character until the character before the last one:

- get current char and next chars decimal equivalents save in variables: currentChar, NextChar
if currentChar < nextChar {
	convertedNum = nextChar - currentChar
} else {
	convertedNum += nextChar + currentChar
}
increment by 2

return result


*/