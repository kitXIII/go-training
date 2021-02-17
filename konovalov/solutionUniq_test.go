package konovalov

import (
	"testing"
)

type uniqTestPairs struct {
	values   []int
	expected int
}

var uniqTests = []uniqTestPairs{
	{[]int{1, 2, 3, 4, 5}, 5},
	{[]int{}, 0},
	{[]int{1, 1, 1, 1, 1}, 1},
	{[]int{1, 1, 2, 2, 3}, 3},
}

func TestSolutionUniq(t *testing.T) {
	for _, pair := range uniqTests {
		result := SolutionUniq(pair.values)
		if result != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", result,
			)
		}
	}
}
