package konovalov

import (
	"math"
)

// SolutionBinaryGap calculates the maximum number of zeros in binary representation
func SolutionBinaryGap(number int) int {
	if number == math.MinInt64 {
		return 63
	}

	num := abs(number)

	steps := 64
	if number < 0 {
		steps = 63
	}

	iter := getIterator()
	max := 0

	for i := 0; i < steps; i++ {
		max = iter(num & 1)
		num = num >> 1
	}

	return max
}

func getIterator() func(int) int {
	max, current := 0, 0

	return func(lastBit int) int {
		if lastBit == 0 {
			current++
		} else if max > current {
			current = 0
		} else {
			max, current = current, 0
		}

		if max > current {
			return max
		}

		return current
	}
}

func abs(num int) int {
	if num >= 0 {
		return num
	}

	return -num
}
