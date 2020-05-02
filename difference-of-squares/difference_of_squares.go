package diffsquares

//SquareOfSum takes as input the first N natural numbers and return the squar root of their sum
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

//SumOfSquares takes as input the first N natural numbers and return the sum of the squart root of each number
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

//Difference should return diference between two numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
