package main

import (
	"fmt"
	"strconv"
	"math"
	"strings"
)

func Encode(text string) (tripledBinary string) {
  completeBinary := ""
	for _, char := range text {
    asBinary := convertDecimalToBinary(int(char))
		completeBinary += asBinary
	}
  
	for _, binaryDigit := range completeBinary {
		tripledBinary += strings.Repeat(string(binaryDigit), 3)
	}
  return
}

func Decode(bits string) (outputStr string) {
	triplesList := []string{}

	for index := 0; index < len(bits); index += 3 {
		tempStr := string(bits[index]) + string(bits[index + 1]) + string(bits[index + 2])
		triplesList = append(triplesList, returnMostCommonOutOfThree(tempStr))
	}
  fullString := strings.Split(strings.Join(triplesList, ""), "")

	for index := 0; index < len(fullString); index += 8 {
    eightBits := strings.Join(fullString[index:index+8], "")
		outputStr += string(rune(convertBinaryToDecimal(eightBits)))
	}
  return
}

func returnMostCommonOutOfThree(str string) string {
	if string(str[0]) == string(str[2]) || string(str[0]) == string(str[1]) {
		return string(str[0])
	} else if string(str[1]) == string(str[2]) {
		return string(str[1])
	}
	return ""
}

func convertDecimalToBinary(num int) (binaryStr string) {
	intermediateResult := num

	for intermediateResult >= 1 {
		temp := strconv.Itoa(intermediateResult % 2)
		binaryStr = temp + binaryStr
		intermediateResult = intermediateResult / 2
	}
	return fmt.Sprintf("%08s",binaryStr)
}

func convertBinaryToDecimal(binaryStr string) int {
	count := len(binaryStr) - 1
	var convertedNum float64 = 0
	for _, digit := range binaryStr {
		digitAsFloat, _ := strconv.ParseFloat(string(digit), 64)
		convertedNum += (digitAsFloat * math.Pow(2.0, float64(count)))
	  count--
	}
	return int(convertedNum)
}

func main() {
	fmt.Println(Decode("100111111000111001000010000111111000000111001111000111110110111000010111")) // "hey"
	fmt.Println(Encode("hey")) //"100111111000111001000010000111111000000111001111000111110110111000010111"
}