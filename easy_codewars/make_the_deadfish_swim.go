package main

import (
	"fmt"
	"strings"
)

func Parse(data string) []int {

	resultSlice := []int{}
	value := 0
	dataList := strings.Split(data, "")

	for _, el := range dataList {
		switch el {
		case "i":
			value += 1
		case "d":
			value -= 1
		case "s":
			value *= value
		case "o":
			resultSlice = append(resultSlice, value)
		default:
			continue
		}
	}
	return resultSlice
}

func main() {
	fmt.Println(Parse("iiisdoso"))
}