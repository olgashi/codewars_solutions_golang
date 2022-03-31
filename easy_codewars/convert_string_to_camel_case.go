package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	inputStringSlice := regexp.MustCompile("[_-]").Split(s, -1)
	for index := 0; index < len(inputStringSlice); index++ {
		if index == 0 {
			continue
		}
		inputStringSlice[index] = strings.ToUpper(inputStringSlice[index][:1]) + inputStringSlice[index][1:]
	}
	
	return strings.Join(inputStringSlice, "")
}

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior")) //"theStealthWarrior"
	fmt.Println(ToCamelCase("The_Stealth_Warrior")) //"TheStealthWarrior"
	fmt.Println(ToCamelCase("HelloWorld")) // "HelloWold"
	fmt.Println(ToCamelCase("helloWorld")) // "helloWorld"
	fmt.Println(ToCamelCase("")) // ""
	fmt.Println(ToCamelCase("hello_hello-hello_World")) // "helloHelloHelloWorld"
	fmt.Println(ToCamelCase("what_a-beautiful_day!")) // "whatABeautifulDay!"
}

/*
- convert dash or underscore delimited words into camel casing
- first word in the output should be capitalized ONLY if the original word was capitalized
- any other word (that comes after - or _ delimiter should be capitalized)

Questions:
Should we expect any other types of delimiters besides -_? No
Can there be numbers? Can an input string start with a number? No

Should we expect an empty string? Yes, return empty string
Should we expect a string containing only - and _? No
Should we expect any special characters (punctuation) ? Yes, keep as is


"the-stealth-warrior" gets converted to "the-stealth-warrior"
"The_Stealth_Warrior" gets converted to "TheStealthWarrior"

DS:
-slice
// the-stealth-warrior

Algorithm:
- split input string on _ or - , save to variable inputStringSlice

- iterate over the inputStringSlice, for each element:
- if index == 0 {
	continue
}
  upcase the first letter in the current word and save back in the same slot in inputStringSlice

return inputStringSlice joined on empty space
*/