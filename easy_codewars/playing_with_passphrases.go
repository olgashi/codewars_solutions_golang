package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PlayPass(s string, n int) string {
	outputStr := ""

	for index := 0; index < len(s); index++ {
		char := s[index]
		charToAppend := ""

		if char >= 65 && char <= 90 {
			char = s[index] + byte(n)
			if char > 90 {
				char -= 26
			}

			if index%2 == 0 {
				charToAppend = strings.ToUpper(string(char))
			} else {
				charToAppend = strings.ToLower(string(char))
			}

		} else if char >= 48 && char <= 57 {
			convertedNum, _ := strconv.Atoi(string(char))
			complementTo9 := 9 - convertedNum
			charToAppend = strconv.Itoa(complementTo9)
		} else {
			charToAppend = string(char)
		}

		outputStr = charToAppend + outputStr
	}
	return outputStr
}

func main() {
	fmt.Println(PlayPass("BORN IN 2015!", 1))    // CPSO JO 7984! => CpSo jO 7984! => !4897 Oj oSpC
	fmt.Println(PlayPass("HELLO WORLD? 123", 2)) // JGNNQ YQTNF?0 876 => JgNnQ YqTnF? 876 => 678 ?FnTqY QnNgJ
	fmt.Println(PlayPass("I LOVE YOU!!!", 0))    // !!!uOy eVoL I
	fmt.Println(PlayPass("AAABBCCY", 1))         // zDdCcBbB
}
