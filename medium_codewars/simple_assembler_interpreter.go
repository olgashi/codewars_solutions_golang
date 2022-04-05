package main

import (
	"fmt"
	"strings"
	"strconv"
)

func SimpleAssembler(program []string) map[string]int {
  registerMap := make(map[string]int)
  index := 0
	for index < len(program) {
		instructionArr := strings.Split(program[index], " ")
		instruction, val1 := instructionArr[0], instructionArr[1]
    val2 := ""
		if len(instructionArr) > 2 {
			val2 = instructionArr[2]
		}
		
		switch instruction {
		case "mov":
			valueToSet := 0
			asciiValue := []rune(val2)[0]
			if asciiValue >= 97 && asciiValue <= 122 {
				valueToSet = registerMap[val2]
			} else {
				valueToSet, _ = strconv.Atoi(val2)
			}
			registerMap[val1] = valueToSet
		case "inc":
			registerMap[val1]++
		case "dec":
			registerMap[val1]--
		case "jnz":
			asciiValue := []rune(val1)[0]
			val2ToInt, _ := strconv.Atoi(val2)
			if asciiValue >= 97 && asciiValue <= 122 && registerMap[val1] != 0 || 
			((asciiValue >= 48 && asciiValue <= 57) && val2ToInt != 0) { 
				index = index + val2ToInt
				continue
			}
		}
		index++
	}
	return registerMap
}

func main() {
	fmt.Println(SimpleAssembler([]string{"mov a 1", "mov b 1", "mov c 0", "mov d 26", 
	"jnz c 2", "jnz 1 5", "mov c 7", "inc d", "dec c", "jnz c -2", "mov c a", "inc a ",
	"dec b", "jnz b -2", "mov b c", "dec d", "jnz d -6", "mov c 18", "mov d 11", "inc a", 
	"dec d", "jnz d -2", "dec c", "jnz c -5"})) //map[a:318009 b:196418 c:0 d:0]
	fmt.Println(SimpleAssembler([]string{"mov a 5", "inc a", "dec a", "dec a", "jnz a -1", "inc a"})) // {"a": 1}
	fmt.Println(SimpleAssembler([]string{"mov a -10", "mov b a", "inc a", "dec b", "jnz a -2"})) //{"a": 0, "b": -20}
	fmt.Println(SimpleAssembler([]string{"mov a 1", "mov b 1", "mov c 1", "dec c", "dec a","inc b", "dec b", "jnz -1"})) //{"a":0, "b": 0, "c":0}
	fmt.Println(SimpleAssembler([]string{"mov a 10", "mov b a", "dec a", "dec a", "jnz -2"})) //{"a": 0, "b":10}

}
