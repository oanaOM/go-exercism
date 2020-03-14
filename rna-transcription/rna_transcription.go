/*
* Given a DNA strand, its transcribed RNA strand is formed by replacing each nucleotide with its complement:
* G -> C
* C -> G
* T -> A
* A -> U
 */

package strand

import (
	"strings"
)

//ToRNA func takes a given DNA strand as input and return its RNA complement (per RNA transcription
func ToRNA(dna string) string {
	return strings.NewReplacer("G", "C", "C", "G", "T", "A", "A", "U").Replace(dna)
}
