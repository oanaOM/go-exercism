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
	"strings"
)

//IsIsogram checks if a word/phrase is isogram
func IsIsogram(phrase string) bool {
	//convert to lower case
	phrase = strings.ToLower(phrase)
	//replace all hyphen
	phrase = strings.ReplaceAll(phrase, "-", "")
	//replace all spaces
	phrase = strings.ReplaceAll(phrase, " ", "")

	if len(phrase) != 1 && len(phrase) != 0 {
		for i := 0; i < len(phrase); i++ {
			for j := i + 1; j < len(phrase); j++ {
				if phrase[i] == phrase[j] {
					return false
				}
			}
		}
	}

	return true

}
