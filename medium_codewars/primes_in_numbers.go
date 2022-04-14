package main

import (
	"fmt"
	"math/big"
)

func PrimeFactors(n int) string {
	num := int64(n)
	var factor int64 = 2
	factorList := []int64{}

	for num > 1 {
		if (num % factor == 0) || num == 1 {
			factorList = append(factorList, factor)
			num = num/factor
		} else {
			factor++
		}
	}

	outputStr := ""
	primes := make(map[int64]int)
	keys := []int64{}
	for index := 0; index < len(factorList); index++ {
		_, ok := primes[factorList[index]]
		if !ok {
			keys = append(keys, factorList[index])
		}
		if ok || big.NewInt(factorList[index]).ProbablyPrime(0) {
			primes[factorList[index]]++
		}
	}

	for _, key := range keys {
		val := primes[int64(key)]
		if val == 1 {
			outputStr += "(" + fmt.Sprint(key) + ")"
		} else {
			outputStr += "(" + fmt.Sprint(key) + "**" + fmt.Sprint(val) + ")"
		}
	}
	return outputStr
}

func main() {
  fmt.Println(PrimeFactors(86240)) // "(2**5)(5)(7**2)(11)"
  fmt.Println(PrimeFactors(7775460)) // "(2**2)(3**3)(5)(7)(11**2)(17)"
  fmt.Println(PrimeFactors(79340)) // "(2**2)(5)(3967)"
}