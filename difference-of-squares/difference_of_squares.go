// Package diffsquares provided utility functions for square and sum operations
package diffsquares

// squareOfSums maps n to square of sum of the numbers up to n
var squareOfSums = make(map[int]int)

// sumOfSquares maps n to sum of squares of the numbers up to n
var sumOfSquares = make(map[int]int)

// SquareOfSum sums all the numbers upto n and squares the sum
func SquareOfSum(n int) int {
	res, ok := squareOfSums[n]
	if ok {
		return res
	}
	res = (n * (n + 1) / 2) * (n * (n + 1) / 2)
	squareOfSums[n] = res
	return res
}

// SumOfSquares squares all the numbers upto n and sums them
func SumOfSquares(n int) int {
	res, ok := sumOfSquares[n]
	if ok {
		return res
	}
	res = n * (n + 1) * (2*n + 1) / 6
	sumOfSquares[n] = res
	return res
}

// Difference returns the difference between square of sum and sum of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
