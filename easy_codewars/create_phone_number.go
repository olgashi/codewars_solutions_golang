package main

import (
	"fmt"
)

func CreatePhoneNumber(nums [10]uint) string {
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", nums[0], nums[1], nums[2], nums[3], nums[4], nums[5], nums[6], nums[7], nums[8], nums[9])
}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})) // (123) 456-7890
	fmt.Println(CreatePhoneNumber([10]uint{7, 2, 3, 4, 1, 6, 7, 9, 9, 8})) // (723) 416-7998
	fmt.Println(CreatePhoneNumber([10]uint{2, 2, 2, 4, 0, 6, 9, 0, 9, 9})) // (222) 406-9099
}