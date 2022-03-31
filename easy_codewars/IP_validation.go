package main

import (
	"fmt"
	"strings"
	"strconv"
	"regexp"
)


func Is_valid_ip(ip string) bool {
	ipOctetsList := strings.Split(ip, ".")


	if len(ipOctetsList) != 4 {
		return false
	}

	for _, el := range ipOctetsList {
		if matched, _ := regexp.Match("[a-z]", []byte(el)); 
		matched ||
		(el != "0" && string(el[0]) == "0") || 
		(strings.Contains(el, " ")) {
			return false
		}
		
		octetAsInt, _ := strconv.Atoi(el)
		
		if octetAsInt < 0 || octetAsInt > 255 {
			return false
		} 
	}

	return true
}


func main() {
	fmt.Println(Is_valid_ip("1.2.3.4")) // valid
	fmt.Println(Is_valid_ip("123.45.67.89")) // valid
	fmt.Println(Is_valid_ip("1.2.3")) // false
	fmt.Println(Is_valid_ip("1.2.3.4.5")) // false
	fmt.Println(Is_valid_ip("123.456.78.90")) // false
	fmt.Println(Is_valid_ip("123.045.067.089")) // false
	fmt.Println(Is_valid_ip("345.1.2.3")) // false
	fmt.Println(Is_valid_ip("123.256.256.256")) // false
	fmt.Println(Is_valid_ip("123.255.255.255")) // true
	fmt.Println(Is_valid_ip("1.abc.255.255")) // false
}

/*
PEDAC

Input:
string, consists of four dot separated octets

Output:
boolean, indicating if IP is valid or not

Rules:
inputs are guaranteed to be a single string

valid ip should:
consits of 4 octets
each octet is an integer between 0 and 255
cannot start with a zero (unless it is a zero)

Examples:

DS 
  - slice to process each octet

Algorithms
  - split the input string on ., save into ipOctetsList variable

	if lenght of ipOctetsList is not == 4 
	return false

loop over ipOctetsList, for each element:
  - if octet starts with 0 but not equal to 0
	  return false
	- if octet converted to int is NOT  >= 0  and <= 255
	  return false


*/