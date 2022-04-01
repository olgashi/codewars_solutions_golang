package main

import (
	"fmt"
	"strings"
	"math"
)

func FirstNonRepeating(str string) string {
  letterCountMap := make(map[string]map[string]int)

  for index, el := range str {
		letter := strings.ToLower(string(el))
		if _, exists := letterCountMap[letter]; !exists {
			letterCountMap[letter] = make(map[string]int)
		}

		letterCountMap[letter]["index"] = index
		letterCountMap[letter]["count"]++
	}
  firstOccuringUniqueLetter := ""
	smallestIndex := math.Inf(1)

	for _, v := range letterCountMap {
		if smallestIndex > float64(v["index"]) && v["count"] == 1{
			smallestIndex = float64(v["index"])
			firstOccuringUniqueLetter = string(str[int(smallestIndex)])
		}
	}

	return firstOccuringUniqueLetter
}

func main() {
  fmt.Println(FirstNonRepeating("sTress")) // 'T'
	fmt.Println(FirstNonRepeating("stress")) // 't'
	fmt.Println(FirstNonRepeating("sttrerress")) // ""
	fmt.Println(FirstNonRepeating("abcd")) // "a"
	fmt.Println(FirstNonRepeating("aaabbccdde")) // "e"
}

/* PEDAC

Problem
_______

Input: string
Outputs: first non repeating character

Rules:
- first non repeating charater - it is the first character that is not repeated anywhere in the string 
(unique, appears only 1 time)

- upper and lower case letters are considered the same character
- the function should return the correct case for the initial letter

sTreSS, should return 'T' not 't'

- if string contains all repeating characters, return an empty string - ""


Examples
_______

	fmt.Println(FirstNonRepeating("sTress")) // 'T'
	fmt.Println(FirstNonRepeating("stress")) // 't'
	fmt.Println(FirstNonRepeating("sttrerress")) // ""
	fmt.Println(FirstNonRepeating("abcd")) // "a"
DS
__
 map "letter" : {count: int, index: 0}

Algorithm
_________
create variable - letterCountMap

loop over the whole input string, for each letter:
		add to map: "letter in lower cased form": {count: int, index:int}

create variable firOccuringUniqueLetter = ""
loop over map, find the letter with the lowest index
  access the letter in the original string at the index, set it to firOccuringUniqueLetter

return firOccuringUniqueLetter


Code
____



*/
