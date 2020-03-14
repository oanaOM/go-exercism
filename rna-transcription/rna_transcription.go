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
	dnaR := []rune(dna)
	//initiliaze our RNA
	var rna string
	//initliaze a nucleotides DNA vs RNA map
	m := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}
	for i := range dnaR {
		for k, v := range m {
			if string(dnaR[i]) == k {
				rna += v
			}
		}
	}
	return rna
}
