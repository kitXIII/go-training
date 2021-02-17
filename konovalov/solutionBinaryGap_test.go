package konovalov

import (
	"math"
	"testing"
)

type binaryTestPair struct {
	value    int
	expected int
}

var binaryTests = []binaryTestPair{
	{-1, 62},                 // 1000000000000000000000000000000000000000000000000000000000000001
	{0, 64},                  // 0000000000000000000000000000000000000000000000000000000000000000
	{math.MinInt64, 63},      // 1000000000000000000000000000000000000000000000000000000000000000
	{math.MaxInt64, 1},       // 0111111111111111111111111111111111111111111111111111111111111111
	{144115188075855871, 7},  // 0000000111111111111111111111111111111111111111111111111111111111
	{-144115188075855871, 6}, // 1000000111111111111111111111111111111111111111111111111111111111
}

func TestSolutionBinaryGap(t *testing.T) {
	for _, pair := range binaryTests {
		result := SolutionBinaryGap(pair.value)
		if result != pair.expected {
			t.Error(
				"For", pair.value,
				"expected", pair.expected,
				"got", result,
			)
		}
	}
}