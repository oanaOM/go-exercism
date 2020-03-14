/*
* Given a DNA strand, its transcribed RNA strand is formed by replacing each nucleotide with its complement:
* G -> C
* C -> G
* T -> A
* A -> U
 */

package strand

//ToRNA func takes a given DNA strand as input and return its RNA complement (per RNA transcription
func ToRNA(dna string) string {
	//convert our string into a rune
	dnaR := []byte(dna)
	//initiliaze our RNA
	//var rna string
	//initliaze a nucleotides DNA vs RNA map
	nucleotides := map[byte]byte{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
	for i, j := range dnaR {
		dnaR[i] = nucleotides[j]
	}
	return string(dnaR)
}
