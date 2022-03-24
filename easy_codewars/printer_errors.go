package main

import (
	"fmt"
)

//PEDAC

/*

- a...m represent different colors, where each unique letter respresents unique color
- any letter staring with n is not a color

- good control string: only valid colors appear (letters representing colors) example: aaabbbbhaijjjm
- bad control string: there will be letters that are not in the range of a..m example: aaaxbbbbyyhwawiwjjjwwm

- objective return a string which is a rational, for example 2/14
  where numeratior - 2 is number of "error" letters appearing in the input string
	and denominator - 14 is the total length of the control string

- do not reduce fraction to a simpler expression

- input string has a length greater or equal to 1
- input string contains only letters from a...z


// Input: string representing a control string
// Output: string representing a rational of number of error letters over input string length (2/14)


// Questions:
- Can there be upper case letters? Or other characters? No

// Examples:
done

Algorithm:
	fmt.Println(PrinterError("lllllxxxx")) // 4/9

- create variable errorLetterCount of type int
loop over the input string, for each leeter in the string:
  if letter is greater than m
	  increment errorLetterCount

return string(errorLetterCount) + "/" + len(s)

*/

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