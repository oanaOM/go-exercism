/*Package luhn implements the Luhn algorithm
* Step 1: Double every second digit starting from the right
* Step 2: Sum all of the digits
* Step 3: If sum is evenly devisible by 10, then the number is valid.
**/
package luhn

import (
	"strconv"
	"strings"
)

// Valid function checks if a string is valid using the Luhn algorithm
func Valid(input string) bool {

	var sum int
	//initiate a bool that will be used to alternate the element position
	second := false
	//remove all spaces
	input = strings.ReplaceAll(input, " ", "")
	//we can't validate single values
	if len(input) <= 1 {
		return false
	}
	inputR := []rune(input)

	for i := len(inputR) - 1; i >= 0; i-- {
		//convert string to int
		cn, err := strconv.Atoi(string(inputR[i]))
		if err != nil {
			return false
		}
		if second {
			cn *= 2
			if cn > 9 {
				cn -= 9
			}
		}
		//switch position
		second = !second
		sum += cn
	}
	return sum%10 == 0
}
