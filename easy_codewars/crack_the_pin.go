package main

import (
	"fmt"
	"crypto/md5"
	"strconv"

)

func Crack(hash string) string {
	for count := 0; count <= 99999; count++ {
		possiblePin := fmt.Sprintf("%05s", strconv.Itoa(count))
		data := []byte(possiblePin)
		possibleHash := fmt.Sprintf("%x", md5.Sum(data))
		if possibleHash == hash {
			return possiblePin
		}
	} 

  return ""
}

func main() {
	fmt.Println(Crack("827ccb0eea8a706c4c34a16891f84e7b")) // 12345
}