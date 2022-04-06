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