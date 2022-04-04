package main

import (
	"fmt"
	"strings"
)

func solve(str string) int {
	var abc map[string]int = map[string]int{"a":1, "b":2, "c":3, "d":4, "e":5, "f":6, "g":7, "h":8, "i":9, "j":10, "k":11, "l":12, "m":13, "n":14, "o":15, "p":16, "q":17, "r":18, "s":19, "t":20, "u":21, "v":22, "w":23, "x":24, "y":25, "z":26}
  substrCount := 0
	finalMax := 0
	for index, letter := range strings.ToLower(str) {
		letterStr := string(letter)
		if letterStr == "a" || letterStr == "e" || letterStr == "i" || 
		   letterStr == "o" || letterStr == "u" {
			if finalMax < substrCount {
				finalMax = substrCount
			}
			substrCount = 0
		} else if index == len(str) - 1 {
			substrCount += abc[letterStr]
			if substrCount > finalMax {
				finalMax = substrCount
			}
		} else {
			substrCount += abc[letterStr]
		}
	}
	
	return finalMax
}
 
func main() {
  // fmt.Println(solve(""))
  fmt.Println(solve("abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes")) // 143
  fmt.Println(solve("zodiacs")) //26
  fmt.Println(solve("strength"))//57
  fmt.Println(solve("zoo")) //26
  fmt.Println(solve("bcd")) //9
}