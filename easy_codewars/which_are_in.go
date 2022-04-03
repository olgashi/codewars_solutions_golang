package main

import (
	"fmt"
	"strings"
	"sort"
)

func InArray(array1 []string, array2 []string) []string {
	result := []string{}

	for indexArr1 := 0; indexArr1 < len(array1); indexArr1++ {
		for indexArr2 := 0; indexArr2 < len(array2); indexArr2++ {
			if strings.Contains(array2[indexArr2], array1[indexArr1]) &&
			   !sliceContainsString(result, array1[indexArr1]) {
				result = append(result, array1[indexArr1])
				}
			}
		}
		sort.Strings(result)

	return result
}

func sliceContainsString(s []string, val string) bool {
	for _, str := range s {
		if str == val {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(InArray([]string{"live", "arp", "strong"}, []string{"lively", "alive", "harp", "sharp", "armstrong"})) //[arp live strong]
}