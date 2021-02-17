package konovalov

// SolutionUniq conuts uniq values in the given array of integers
func SolutionUniq(arr []int) int {
	acc := make(map[int]bool)

	for _, v := range arr {
		acc[v] = true
	}

	return len(acc)
}
