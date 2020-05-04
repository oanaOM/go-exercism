/*Package luhn implements the Luhn algorithm
* Step 1: Double every second digit starting from the right
* Step 2: Sum all of the digits
* Step 3: If sum is evenly devisible by 10, then the number is valid.
**/
package luhn

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Valid function checks if a string is valid using the  Luhn algorithm
func Valid(input string) bool {

	var sum, j int
	//remove all spaces
	input = strings.Replace(input, " ", "", -1)
	//we can't validate single values
	if len(input) == 0 || len(input) == 1 {
		return false
	}
	inputR := []rune(input)
	for i := len(inputR) - 1; i >= 0; i-- {
		cn, err := strconv.Atoi(string(inputR[i]))
		if err != nil {
			fmt.Printf("Invalid SIN.")
		}
		//if a string contains a non-numberic element, is invalid
		if !unicode.IsNumber(inputR[i]) {
			return false
		}
		//use j to read from the last index
		j++
		if j%2 == 0 {
			cnDouble := cn * 2
			if cnDouble > 9 {
				cn = cnDouble - 9
			} else {
				cn = cnDouble
			}
		}
		sum += cn
	}
	return sum%10 == 0
}
