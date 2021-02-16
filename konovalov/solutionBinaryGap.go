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

	max := 0
	current := 0

	for i := 0; i < 63; i++ {
		max, current = handleLastBit(num&1, max, current)
		num = num >> 1
	}

	if max > current {
		return max
	}

	return current
}

func handleLastBit(lastBit int, max int, current int) (int, int) {
	if lastBit == 0 {
		return max, current + 1
	}

	if max > current {
		return max, 0
	}

	return current, 0
}

func abs(num int) int {
	if num >= 0 {
		return num
	}

	return -num
}
