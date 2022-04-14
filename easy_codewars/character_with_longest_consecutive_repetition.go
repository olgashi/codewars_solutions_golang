package main

import (
	"fmt"
)

type Result struct {
	C rune // character
	L int  // count
}

func LongestRepetition(text string) (letterCount Result) {
  if len(text) == 0 {
		return
	}

  letterCount.C = rune(text[0])
	letterCount.L = 1
	previousLetter := string(text[0])
	currentLetterCount := 1

	for index := 1; index < len(text); index++ {
		currentLetter := string(text[index])

		if previousLetter == currentLetter {
			currentLetterCount++
			
			if (index == len(text) - 1) && currentLetterCount > letterCount.L {
				letterCount.C = rune(text[index])
				letterCount.L = currentLetterCount
			}
			} else {
				if currentLetterCount > letterCount.L {
					letterCount.C = rune(text[index-1])
					letterCount.L = currentLetterCount
				}
				currentLetterCount = 1
			}
			
			previousLetter = string(text[index])
	}
	return letterCount
}

func main() {
	fmt.Println(LongestRepetition("aaaabb")) //: a: 4, 
	fmt.Println(LongestRepetition("bbbaaabcccaaaa"))
	fmt.Println(LongestRepetition("cbdeuuu900"))
	fmt.Println(LongestRepetition("abbbbb"))
	fmt.Println(LongestRepetition("aabb"))
	fmt.Println(LongestRepetition("ba"))
	fmt.Println(LongestRepetition(""))
}
