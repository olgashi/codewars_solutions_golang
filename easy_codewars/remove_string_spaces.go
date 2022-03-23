package main
import (
	"fmt"
	"strings"
)

func NoSpace(word string) string {
	// var strNoSpaces = ""

	// for _, char := range word {
	// 	if string(char) != string(" ") {
	// 		strNoSpaces += string(char)
	// 	}
	// }
	// return strNoSpaces
	
	// or
	return strings.ReplaceAll(word, " ", "")
}


func main() {
	fmt.Println(NoSpace("aaa ddd s rrrrr"))

}