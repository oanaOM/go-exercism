/*Package scrabble compute the Scrabble score for that word.
Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
*/
package scrabble

import (
	"strings"
)

//LetterCount counts how many letters of same type are in a word
func LetterCount(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range []rune(strings.ToLower(s)) {
		m[string(w)]++
	}
	return m
}

//Score takes a word as parameter and calculates it's scrore
func Score(word string) int {
	//initializing our score
	var score int
	//create a map having the letter point as key and corresponding letter as value
	scoreMap := map[int]string{
		1:  "AEIOULNRST",
		2:  "DG",
		3:  "BCMP",
		4:  "FHVWY",
		5:  "K",
		8:  "JX",
		10: "QZ",
	}
	for _, letter := range strings.Split(strings.ToLower(word), "") {
		for k, v := range scoreMap {
			if strings.Contains(strings.ToLower(v), letter) {
				score += k
			}
		}

	}
	return score
}
