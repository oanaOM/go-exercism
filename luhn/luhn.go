/*Package luhn implements the Luhn algorithm
* Step 1: Double every second digit starting from the right
* Step 2: Sum all of the digits
* Step 3: If sum is evenly devisible by 10, then the number is valid.
**/
package luhn

import (
	"strings"
	"unicode"
)

// Valid function uses Luhn algorithm to check if a string is valid or not
func Valid(input string) bool {

	var sum int
	//remove all spaces
	input = strings.ReplaceAll(input, " ", "")
	//we can't validate single values
	if len(input) <= 1 {
		return false
	}
	//initiate a bool that will be used to alternate the element position
	second := len(input)%2 == 0
	for _, r := range input {
		if !unicode.IsDigit(r) {
			return false
		}
		cn := int(r - '0')
		if second {
			cn *= 2
			if cn > 9 {
				cn -= 9
			}
		}
		sum += cn
		second = !second
	}
	return sum%10 == 0
}
