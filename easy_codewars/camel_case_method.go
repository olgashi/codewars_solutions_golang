package main

import (
	"fmt"
	"strings"
)

func CamelCase(s string) (outputWord string) {
  if len(s) == 0 {
    return ""
  }
  for _, word := range strings.Split(strings.Trim(s, " "), " ") {
    outputWord += strings.ToUpper(string(word[0])) + word[1:]
  }
  return outputWord
}

func main() {
  fmt.Println(CamelCase("")) // ""
  fmt.Println(CamelCase("camel case word")) //CamelCaseWord
}