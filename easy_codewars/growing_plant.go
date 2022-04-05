package main

import (
	"fmt"
)

func GrowingPlant(upSpeed, downSpeed, desiredHeight int) (numDays int) {
	growingPlantHeight := 0

	for {
		numDays++
		if growingPlantHeight += upSpeed; growingPlantHeight >= desiredHeight {
		  return numDays
		}

		growingPlantHeight -= downSpeed
	}
}


func main() {
	fmt.Println(GrowingPlant(100, 10, 910)) // 10
	fmt.Println(GrowingPlant(10, 9, 4)) // 1
	fmt.Println(GrowingPlant(5, 2, 5)) // 1
	fmt.Println(GrowingPlant(5, 2, 6)) // 2
	fmt.Println(GrowingPlant(50, 10, 150)) // 4
	fmt.Println(GrowingPlant(5, 2, 17)) // 5
}