/*Package isogram has a function that checks is a word/phrase is isogram
An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.
* Examples of isograms:
	lumberjacks
	background
	downstream
	six-year-old

non-isogram words: isograms
*/
package isogram

import (
	"unicode"
)

//IsIsogram checks if a word/phrase is isogram
func IsIsogram(phrase string) bool {

	runes := map[rune]bool{}

	for _, r := range phrase {

		r = unicode.ToLower(r)
		if runes[r] && unicode.IsLetter(r) {
			return false
		}

		runes[r] = true
	}

	return true

}
