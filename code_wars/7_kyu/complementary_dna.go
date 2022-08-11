// Your function receives one side of the DNA; you need to return the other complementary side. DNA strand is never empty or there is no DNA at all.

// "ATTGC" --> "TAACG"
// "GTAT" --> "CATA"

package main

import (
	"fmt"
)

func DNAStrand(dna string) string {
	dnaComplements := map[string]string{"A": "T", "T": "A", "G": "C", "C": "G"}
	var complementStrand string

	for _, value := range dna {
		complementStrand += dnaComplements[string(value)]
	}

	return complementStrand
}

// func DNAStrand(dna string) string {
// 	replacer := strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G")
// 	return (replacer.Replace(dna))
// }

func main() {
	fmt.Println(DNAStrand("AAAA"))
	fmt.Println(DNAStrand("ATTGC"))
	fmt.Println(DNAStrand("GTAT"))
}
