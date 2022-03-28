package main

import (
	"fmt"
)

func FindOutlier(integers []int) int {
	index := 1
	
	for index < len(integers) {
		isNum1Odd := integers[index - 1] % 2 == 0
		isNum2Odd := integers[index] % 2 == 0
		isNum3Odd := integers[index + 1] % 2 == 0

		if (isNum1Odd && !isNum2Odd && !isNum3Odd) || (!isNum1Odd && isNum2Odd && isNum3Odd) {
			return integers[index - 1]
		} else if (!isNum1Odd && isNum2Odd && !isNum3Odd) || (isNum1Odd && !isNum2Odd && isNum3Odd) {
			return integers[index]
		} else if (isNum1Odd && isNum2Odd && !isNum3Odd) || (!isNum1Odd && !isNum2Odd && isNum3Odd) {
			return integers[index + 1]
		}

		index++
	}
	return 0
}

func main() {
	fmt.Println(FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
	fmt.Println(FindOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))
}

/*

num1 num2 num3

num1 is odd num2 and num3 are not
num2 is odd num1 and num3 are not
num3 is odd num1 and num2 are not


[2, 4, 0, 100, 4, 11, 2602, 36]
Should return: 11 (the only odd number)

[160, 3, 1719, 19, 11, 13, -21]
Should return: 160 (the only even number)


*/