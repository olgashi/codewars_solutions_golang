package main

import (
	"fmt"
	"strings"
)

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