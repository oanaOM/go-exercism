/*Package grains takes into account a 64 square chessboard and exports two functions:
* Square() which calculates how many grains are on a giving square
*Total() calculates how many grains are in total on board
 */
package grains

import "errors"

//Square calculates the number of grains on a  giving chessboard.
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("Invalid number")
	}
	return 1 << (n - 1), nil
}

//Total calculates how many grains are in total on the chessboard
func Total() uint64 {
	return 1<<64 - 1
}
