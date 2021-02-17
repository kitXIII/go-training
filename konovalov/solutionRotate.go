package konovalov

// SolutionRotate move right round array values
func SolutionRotate(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return arr
	}

	steps := k % len(arr)

	return arrayRotate(arr, steps)
}

func arrayRotate(arr []int, steps int) []int {
	borderIndex := len(arr) - steps
	slice1 := arr[:borderIndex]
	slice2 := arr[borderIndex:]

	return append(slice2, slice1...)
}
