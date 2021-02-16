package konovalov

// SolutionSquareGenerator returns an array of n squares of natural numbers starting from start
func SolutionSquareGenerator(start int, n int) []int {
	results := make([]int, n)
	current := start

	for i := range results {
		results[i] = sqr(current)
		current++
	}

	return results
}

func sqr(num int) int {
	return num * num
}
