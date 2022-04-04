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
