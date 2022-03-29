package main

import (
	"fmt"
	"strings"
)
/*
//PEDAC

// Input: string
// Output: a string, where words with 5 or more letters are reversed

// Requirements:
- All words than are 5 or more letters long are reversed
- Input string will consist of only letters and spaces
- Spaces will be included only when input str contains more than one word

Question:
Will thewordss in the input string always be separated by a single space? Yes

Examples:

spinWords("Hey fellow warriors") // "Hey wollef sroirraw" 
spinWords("This is a test") // "This is a test" 
spinWords("This is another test") // "This is rehtona test"
spinWords("Hello WORLD!") // olleH !DLROW
spinWords("ZOOkeeper") // repeekOOZ

DS:
- slice (list)

// {"This", "is", "another", "test"}
// outputStr: {"This", "is", "rehtona", "test"} // => "This is rehtona test"
Algorithm:

- create a variable to hold output string (slice ds) - outputStr
- split the input string on space (we get a slice)
- iterate over the resulting slice, for each word:
if the word is 5 or more chars long
  reverse the word (either built in function or create a method)

append the resulting word (current word) to outputStr

join outputStr on " ", return

*/
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