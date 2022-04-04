package main

import(
	"fmt"
	"strings"
)
func EncryptThis(text string) (outputStr string) {
	for _, word := range strings.Split(text, " ") {
		for letterIndex := range strings.Split(word, "") {
			if (letterIndex == 0) {
				outputStr += fmt.Sprintf("%v", word[letterIndex])
			} else if (letterIndex == 1) {
				outputStr += string(word[len(word) - 1])
			} else if letterIndex == len(word) - 1 {
				outputStr += string(word[1])
			} else {
				outputStr += string(word[letterIndex])
			}
		}
		outputStr += " "

	}
	return strings.TrimRight(outputStr, " ")
}

func main() {
  fmt.Println(EncryptThis("Hello")) // "72olle"
	fmt.Println(EncryptThis("good")) //"103doo"
	fmt.Println(EncryptThis("hello world")) //"104olle 119drlo"
}