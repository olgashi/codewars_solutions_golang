package main

import (
	"fmt"
)


func endsWith(str, ending string) bool {
	if len(ending) > len(str) {
		return false
	}

	strIndex := len(str) - 1
	for index := len(ending) - 1; index >=0; index-- {
		if strIndex < 0 {
			break
		}
		if (ending[index] != str[strIndex]) {
			return false
		}
		strIndex--
	}
	return true
}

func main() {
	fmt.Println(endsWith("abc", "bc")) // true
	fmt.Println(endsWith("abc", "de")) // false
	fmt.Println(endsWith("HYA!", "!")) // true
	fmt.Println(endsWith("HYA!", " ")) // false
	fmt.Println(endsWith("", "")) // true
	fmt.Println(endsWith(" ", "")) // true
	fmt.Println(endsWith("abc", "c")) // true
	fmt.Println(endsWith("banana", "ana")) // true
	fmt.Println(endsWith("a", "z")) // false
	fmt.Println(endsWith("", "t")) // false
	fmt.Println(endsWith("$a = $b + 1", "+1")) // false
	fmt.Println(endsWith("    ", "   ")) // true
	fmt.Println(endsWith("1oo", "100")) // false
}