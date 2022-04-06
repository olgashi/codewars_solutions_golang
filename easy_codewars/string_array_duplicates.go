package main

import (
	"fmt"
)

func Dup(arr []string) (outputList []string) {

	for _, word := range arr {
		newWord := string(word[0])
		for index := 1; index < len(word); index++ {
			if (word[index] == word[index-1]) {
				continue
			}
			newWord += string(word[index])
		}
		outputList = append(outputList, newWord)
	}
  return
}

func main() {
	fmt.Println(Dup([]string{"abracadabra","allottee","assessee"}))// "abracadabra","alote","asese"
	fmt.Println(Dup([]string{"caaaaatttt","doooogg","dogs"}))// "cat", "dog", "dogs"
	fmt.Println(Dup([]string{"kelless","keenness"}))// "keles","kenes"
}