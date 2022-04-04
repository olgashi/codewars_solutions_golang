package main

import (
	"fmt"
	"strings"
	"strconv"
)

func Revrot(str string, sz int) (output string) {
  if sz <=0 || str == "" || sz > len(str) {
		return ""
	}

	strSlice := strings.Split(str, "")
	strSlice = strings.Split(str, "")[0:len(str) - len(str) % sz]
	
	for index := 0; index < len(strSlice); index += sz {
		currentChunk := strSlice[index:index + sz]

		if isSumOfCubesEven(strings.Join(currentChunk, "")) {
			output += reverseString(strings.Join(currentChunk, ""))
			} else {
      indexToRemove := 0
			indexToInsert := len(currentChunk) - 1
			valueToShift := currentChunk[indexToRemove]

			currentChunk = append(currentChunk[:indexToRemove], currentChunk[indexToRemove+1:]...)
	
			newChunk := make([]string, indexToInsert+1)
			copy(newChunk,currentChunk[:indexToInsert])
			newChunk[indexToInsert] = valueToShift
	
			currentChunk = append(newChunk, currentChunk[indexToInsert:]...)
			output += (strings.Join(currentChunk, ""))
		}
	}

	return output
}

func isSumOfCubesEven(str string) (isEven bool) {
	sum := 0
	for _, num := range str {
		numAsInt, _ := strconv.Atoi(string(num))
		sum += (numAsInt * numAsInt * numAsInt)
	}
	return sum % 2 == 0
}

func reverseString(str string) (reversedStr string) {
	for _, char := range str {
		reversedStr = string(char) + reversedStr
	}
	return
}


func main() {
	fmt.Println(Revrot("123456987654", 6)) // "234561876549"
	fmt.Println(Revrot("123456987653", 6)) // "234561356789"
	fmt.Println(Revrot("66443875", 4)) // "44668753" 
	fmt.Println(Revrot("66443875", 8)) // "64438756"
	fmt.Println(Revrot("664438769", 8)) // "67834466"
	fmt.Println(Revrot("123456779", 8)) // "23456771"
	fmt.Println(Revrot("", 8)) // ""
	fmt.Println(Revrot("123456779", 0)) // "" 
	fmt.Println(Revrot("563000655734469485", 4)) // "0365065073456944")
	fmt.Println(Revrot("1234", 8)) // ""
	fmt.Println(Revrot("1234", -3)) // ""
	fmt.Println(Revrot("123123223", 2)) // 21133222
}
