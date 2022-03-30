package main

import (
	"fmt"
	"strings"
)

/*
Problem

Rule/Requirements:
- symbols A and T are complements of each other
- C and G are also complements

Input: string representing DNA
Output: string, resulting from conversion of symbols to their complement symbols

Questions:
Should I always expect upper cased letters?


Examples

DS
- slice to hold the result of splitting innput string (will contain DNA symbols)
- map containing complement symobols

Algorithm
- create map containing symbols and their complements {"A": "T", "T": "A", "C":"G", "G":"C"}
- split input string on empty space and iterate over resulting slice, for each element
- determine complement symbol for the current element and append to output string
- return output string (will use variable )

Code

*/
func DNAStrand(dna string) (convertedDNASequenceStr string) {
	symbolComplements := map[string]string{"A": "T", "T": "A", "C":"G", "G":"C"}

	for _, symbol := range strings.Split(dna, "") {
		convertedDNASequenceStr += symbolComplements[symbol]
	}

	return
}

func main() {
	fmt.Println(DNAStrand("ATGC")) // "TACG"
	fmt.Println(DNAStrand("GTAT")) // "CATA"
	fmt.Println(DNAStrand("AAAA")) // "TTTT"
	fmt.Println(DNAStrand("GGCT")) // "CCGA"
	fmt.Println(DNAStrand("ATC")) // "TAG"
}