package main

import (
	"fmt"
	"unicode"
	"strconv"
	"strings"
)

func IsSumOfCubes(s string) string {
	numbersList := []string{}
	currentNumber := ""
	sum := 0

	for index, char := range s {
		if char >= 48 && char <= 57 {
			if len(currentNumber) > 0 && len(currentNumber) < 3 {
				currentNumber = currentNumber + string(char)
				} else {
					currentNumber = string(char)
				}
			}
			
			if len(currentNumber) == 3 || (len(currentNumber) > 0 && !unicode.IsDigit(char)) || (len(s) - 1 == index && char >= 48 && char <= 57) {
				if isACubicNumber(currentNumber) {
					num, _ := strconv.Atoi(string(currentNumber))
					intAsStr := strconv.Itoa(num)
					numbersList = append(numbersList, intAsStr)
					sum += num
			}
			currentNumber = ""
		}
	}
	if len(numbersList) == 0 {
		return "Unlucky"
	}
	numbersList = append(numbersList, fmt.Sprintf("%d", sum), "Lucky")
	return strings.Join(numbersList, " ")
}

func isACubicNumber(str string) bool {
	givenNumInteger, _ := strconv.Atoi(str)
  sumOfCubes := 0
	for _, el := range strings.Split(str, "") {
		numInteger, _ := strconv.Atoi(el)
		sumOfCubes += (numInteger * numInteger * numInteger)
	}

	return sumOfCubes == givenNumInteger
}

func main() {
	fmt.Println(IsSumOfCubes("aqdf&0#1xyz!22[153(777.777")) //0 1 153 154 Lucky
	fmt.Println(IsSumOfCubes("&z371 upon 407298a --- dreary, ###100.153 I thought, 9926315strong and weary -127&() 1")) //371 407 153 1 932 Lucky
	fmt.Println(IsSumOfCubes("00 9026315 -827&()")) //0 0 Lucky
	fmt.Println(IsSumOfCubes("sdasdadasda -dasda41sa-")) //Unlucky
}
