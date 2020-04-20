package diffsquares

import (
	"math"
)

type mathExpr = string

const (
	squareOfSum  = mathExpr("SquareOfSum")
	sumOfSquares = mathExpr("SumOfSquares")
)

//MathExpression returns different results depending on the math expression wanted
func MathExpression(expr mathExpr, n int) int {
	var result, sum int
	for i := 1; i <= n; i++ {
		switch expr {
		case squareOfSum:
			sum += i
			result = int(math.Pow(float64(sum), 2))
		case sumOfSquares:
			result += int(math.Pow(float64(i), 2))
		}
	}
	return result
}

//SquareOfSum takes as input the first N natural numbers and return the squar root of their sum
func SquareOfSum(n int) int {
	return MathExpression(squareOfSum, n)
}

//SumOfSquares takes as input the first N natural numbers and return the sum of the squart root of each number
func SumOfSquares(n int) int {
	return MathExpression(sumOfSquares, n)
}

//Difference should return diference between two numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
