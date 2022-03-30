package main

import (
	"fmt"
	"strings"
)

func Solution(str string) (outputSlice []string) {
	
	strSlice := strings.Split(str, "")
  if len(strSlice) % 2 != 0 {
		strSlice = append(strSlice, "_")
	}

	for index := 0; index < len(strSlice) - 1; index +=2 {
		outputSlice = append(outputSlice, strings.Join(strSlice[index:index+2], ""))
	}
  return outputSlice
}

func main() {
	fmt.Println(Solution("abcdefg"))
	
}