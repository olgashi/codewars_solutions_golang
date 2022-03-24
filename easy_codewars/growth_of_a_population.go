package main

import (
	"fmt"
)

func getPopulationEndOfYear(pStart int, percentGrowth float64, additionalPopulation int) float64 {
	return float64(pStart) + (float64(pStart) * (percentGrowth / 100)) + float64(additionalPopulation)
}

func NbYear(p0 int, percent float64, aug int, p int) int {
	var year int = 0
	var populationEndOfYear float64 = float64(p0)

	for populationEndOfYear < float64(p) {
		populationEndOfYear = getPopulationEndOfYear(int(populationEndOfYear), percent, aug)
		year += 1
	}

	return year
}


func main() {
	fmt.Println(NbYear(1000, 2, 50, 1200)) // 3
	fmt.Println(NbYear(100, 30, 10, 500)) // 6 =>100 +30 + 10 => 140, 182+10 => 192, 259.6, 347.48, 461.724, 610.2412
	fmt.Println(NbYear(2000, 10, 600, 4000)) //  3 => 2800, 3680, 4648
}