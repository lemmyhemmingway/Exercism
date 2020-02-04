package diffsquares

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func SquareOfSum(n int) int {
	result := n * (n + 1) / 2
	return result * result
}

func Difference(n int) int {
	n1 := SumOfSquares(n)
	n2 := SquareOfSum(n)
	return n2 - n1
}
