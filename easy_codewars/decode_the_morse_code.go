package main

import (
"strings"
"fmt"
)

func DecodeMorse(morseCode string) string {
	MORSE_CODE := map[string]string{
		".-": "A", "-...": "B", "-.-.":"C", "-..":"D", ".":"E", "..-.":"F", "--.":"G", "....":"H", "..":"I", ".---":"J",
		"-.-":"K", ".-..":"L", "--":"M", "-.":"N", "---":"O", ".--.":"P", "--.-":"Q", ".-.":"R", "...":"S", "-":"T", 
		"..-":"U", "...-":"V",".--":"W", "-..-":"X", "-.--":"Y", "--..":"Z", ".----":"1", "..---":"2", "...--":"3", 
		"....-":"4",".....": "5", "-....":"6", "--...":"7", "---..":"8", "----.":"9", "-----":"0", 
		"...---...": "SOS","-.-.--": "!", "-.--.":"(", "-.--.-":")","-.-.-.":";", "-..-.":"/", "-...-":"=",
    "-....-":"-", ".-.-.-":".",
	}
  var convertedStringSlice []string 

  morseCode = strings.Trim(morseCode, " ")

  morseCodeWordsList := strings.Split(morseCode, "   ")

  for _, word := range morseCodeWordsList {
    word = strings.Trim(word, " ")
    var lettersSlice []string
    for _, letter := range strings.Split(word, " ") {
      lettersSlice = append(lettersSlice, MORSE_CODE[letter])
    }
    convertedWord := strings.Join(lettersSlice, "")
    convertedStringSlice = append(convertedStringSlice, convertedWord)
  }
  return strings.Join(convertedStringSlice, " ")
}

func main() {
  fmt.Println(DecodeMorse(".... . -.--   .--- ..- -.. .")) // HEY JUDE
}