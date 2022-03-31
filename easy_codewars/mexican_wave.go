package main

import (
	"fmt"
	"strings"
)

func wave(words string) []string {
	var waveStrings []string = []string{}
	for index := 0; index < len(words); index++ {
		if string(words[index]) == " " {
			continue
		}

		convertedWord := string(words)[:index] + strings.ToUpper(string(words)[index:index+1]) + string(words)[index+1:]
		waveStrings = append(waveStrings, convertedWord)
	}
	return waveStrings
}


func main() {
	fmt.Println(wave("hello")) //"Hello", "hEllo", "heLlo", "helLo", "hellO"
}