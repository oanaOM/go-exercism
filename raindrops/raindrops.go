/*
* The rules of raindrops are that if a given number:
* has 3 as a factor, add 'Pling' to the result.
* has 5 as a factor, add 'Plang' to the result.
* has 7 as a factor, add 'Plong' to the result.
* does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.
 */

package raindrops

import (
	"strconv"
)

//Convert a number into a string that contains raindrop sounds corresponding to certain potential factors
func Convert(number int) string {
	sounds := [3]string{"Pling", "Plang", "Plong"}
	factor := [3]int{3, 5, 7}
	//initialize an empty raindrops
	var raindrops string

	for v := range sounds {
		if number%factor[v] == 0 {
			raindrops += sounds[v]
		}
	}
	if len(raindrops) == 0 {
		return strconv.Itoa(number)
	}
	return raindrops

}
