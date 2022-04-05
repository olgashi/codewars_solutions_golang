package main

import (
	"fmt"
	"strings"
)

func SpinWords(str string) string {
  var outputStrSlice []string = strings.Split(str, " ")
	for index, word :=range outputStrSlice {
		if len(word) >= 5 {
			 outputStrSlice [index] = reverse(word)
		}
 	}
	 return strings.Join(outputStrSlice, " ")
}

func reverse(str string) (reversedStr string) {
	for _, char := range str {
		reversedStr = string(char) + reversedStr
	}
	return
}

func main() {
	fmt.Println(SpinWords("Hey fellow warriors")) // "Hey wollef sroirraw" 
	fmt.Println(SpinWords("This is a test")) // "This is a test" 
	fmt.Println(SpinWords("This is another test")) // "This is rehtona test"
	fmt.Println(SpinWords("Hello WORLD!")) // olleH !DLROW
	fmt.Println(SpinWords("ZOOkeeper")) // repeekOOZ
}