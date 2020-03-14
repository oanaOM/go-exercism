/*
Calculate the Hamming Distance between two DNA strands
* 2 sequences of equal length
* check the sequence to have exactly the same size
* We read DNA using the letters C,A,G and T
* Sample dataset:
GAGCCTACTAACGGGAT
CATCGTAATGACGGCCT
* Sample output: 7
*/

package hamming

import (
	"fmt"
)

//Distance is the main function
func Distance(a, b string) (int, error) {
	//add a counter
	var counter int
	//convert the strings to runes slice
	ar, br := []rune(a), []rune(b)

	if len(ar) != len(br) {
		return 0, fmt.Errorf("unequal length: %q, %q", a, b)
	}
	for j := range ar {
		if ar[j] != br[j] {
			counter++
		}
	}
	return counter, nil

}
