package main

import (
	"fmt"
	"strings"
)

func LongestConsec(strList []string, k int) string {
	longestString := ""

	if len(strList) == 0 || k > len(strList) || k <= 0 {
		return longestString
	}

	ptr1 := 0
  ptr2 := (ptr1 + k) - 1
	
	for ptr2 < len(strList) {
		tempStr := strings.Join(strList[ptr1:ptr2 + 1], "")

		if len(tempStr) > len(longestString) {
			longestString = tempStr
		}
		ptr1++
		ptr2++
	}
	return longestString
}

func main() {
	fmt.Println(LongestConsec([]string{"tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"}, 2)) // "folingtrashy"
	fmt.Println(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"}, 2)) // "abigailtheta"
	fmt.Println(LongestConsec([]string{"abc", "asd", "asdf", "dsdsds", "awqqwwe", "sasa", "sasa"}, 3)) // "asdfdsdsdsawqqwwe"
	fmt.Println(LongestConsec([]string{"abc", "asd", "asdf", "dsdsds", "sasa", "sasa", "hj", "jklu"}, 4)) // "asdfdsdsdssasasasa"
	fmt.Println(LongestConsec([]string{}, 2)) // ""
	fmt.Println(LongestConsec([]string{"sawq", "dasd", "ddad"}, 3)) // "sawqdasdddad"
	fmt.Println(LongestConsec([]string{"sawq", "dasd", "ddad"}, 4)) // ""
	fmt.Println(LongestConsec([]string{"sawq", "dasd", "ddad"}, -2)) // ""
	fmt.Println(LongestConsec([]string{"sawq", "dasd", "ddad"}, 0)) // ""	
}
