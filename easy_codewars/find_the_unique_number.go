package main

import (
	"fmt"
)

func FindUniq(arr []float32) float32 {
	var uniqueNumsMap = make(map[float32]int)
	var numsArr []float32

	for _, num := range arr {
		uniqueNumsMap[num]++

		if len(numsArr) == 0 || len(numsArr) == 1 && numsArr[0] != num {
			numsArr = append(numsArr, num)
		}
	}

	if (uniqueNumsMap[numsArr[0]] > 1 && uniqueNumsMap[numsArr[1]] == 1) {
		return numsArr[1]

	} else if (uniqueNumsMap[numsArr[0]] == 1 && uniqueNumsMap[numsArr[1]] > 1) {
		return numsArr[0]
	}
	
	return numsArr[0]
}

func main() {
	fmt.Println(FindUniq([]float32{1, 1, 1, 2, 1, 1}))
	fmt.Println(FindUniq([]float32{0.55, 0, 0, 0, 0}))
	
}