package konovalov

import (
	"math"
)

// SolutionBinaryGap calculates the maximum number of zeros in binary representation
func SolutionBinaryGap(number int) int {
	if number == math.MinInt64 {
		return 63
	}

	steps := 64
	if number < 0 {
		steps = 63
	}

	num := abs(number)

	iter := getIterator()
	max := 0

	for i := 0; i < steps; i++ {
		max = iter(num & 1)
		num = num >> 1
	}

	return max
}

// getIterator returns an iterator, that takes 0 or 1 integer number and returns maximum number of continued 0 on each iteration
func getIterator() func(int) int {
	max, current := 0, 0

	return func(nextBitEqualValue int) int {
		if nextBitEqualValue == 0 {
			current++

			if max < current {
				return current
			}

			return max
		}

		if max < current {
			max = current
		}

		current = 0
		return max
	}
}

func abs(num int) int {
	if num >= 0 {
		return num
	}

	return -num
}
