package main

import (
	"fmt"
)

// PEDAC
/*
- each day plant grows by upSpeed meters, 5 <= upSpeed <= 100 
- each night plant decreses by downSpeed meters, 2 <= downSpeed < upSpeed
- initally the plant is 0 meters
- seed is planted at the beginning of day
- objective is to find a day when plant reaches a certain level
- desiredHeight - 4 <= desiredHeight <= 1000, represents treshold
- output - integer representing number of days will take for the plant to reach desiredHeight

// DS
none

// GrowingPlant(5, 2, 6)) // 2
upSpeed = 5 |
downSpeed = 2 |
desiredHeight = 6 |
growingPlantHeight = 0 + 5 - 2 + 5
numDays = 0 + 1 + 1

// Algorithm

create variable growingPlantHeight, set it to 0
create variable numDays = 0

loop while true
  growingPlantHeight += upSpeed
	numDays +=1
	if growingPlantHeight >= desiredHeight
	  return numDays
	
	growingPlantHeight -= downSpeed

*/

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