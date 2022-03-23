package main

func Multiple3And5(number int) int {
	if number <= 0 {
		return 0
	}

	sum := 0

	for currentNum := 1; currentNum < number; currentNum++ {
		if currentNum%3 == 0 {
			sum += currentNum
		} else if currentNum%5 == 0 {
			sum += currentNum
		}
	}
	return sum
}
