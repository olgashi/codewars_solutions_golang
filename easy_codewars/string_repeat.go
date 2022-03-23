package main

// "fmt"
// "strings"

func RepeatStr(repetitions int, value string) string {
	// return strings.Repeat(value, repetitions)
	// or
	var result string = ""
	for index := 0; index < repetitions; index++ {
		result += value
	}
	return result
}

func main() {

}
