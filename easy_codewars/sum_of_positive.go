package main
import (
	"fmt"
)

func PositiveSum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0

	for _, element := range numbers {
		if element > 0 {
			sum += element
		}
	}
	return sum
}

func main() {
	fmt.Println(PositiveSum([]int{1, -4, 7, 12})) // 20
	fmt.Println(PositiveSum([]int{})) // 0
	fmt.Println(PositiveSum([]int{-1, -4, -90})) // 0
	fmt.Println(PositiveSum([]int{1, 4, 12})) // 17
	fmt.Println(PositiveSum([]int{1, 1, 1, 1})) // 4
}