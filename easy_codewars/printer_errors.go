package main

import (
	"fmt"
)

func PrinterError(s string) string {
	var errorLetterCount int

	for _, letter := range s {
    if string(letter) > "m" {
			errorLetterCount++
		}
	}

	return fmt.Sprint(errorLetterCount) + "/" + fmt.Sprint(len(s))
}


func main() {
	fmt.Println(PrinterError("aaabbbbhaijjjm")) // 0/14
	fmt.Println(PrinterError("aaaxbbbbyyhwawiwjjjwwm")) // 8/22
	fmt.Println(PrinterError("xxx")) // 3/3
	fmt.Println(PrinterError("y")) // 1/1
	fmt.Println(PrinterError("lxlxlxxxx")) // 6/9
	fmt.Println(PrinterError("lllllxxxx")) // 4/9
}