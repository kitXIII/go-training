package konovalov

import (
	"testing"
)

type rotateArguments struct {
	array []int
	steps int
}

type rotateTestPairs struct {
	values   rotateArguments
	expected []int
}

var rotateTests = []rotateTestPairs{
	{rotateArguments{[]int{1, 2, 3, 4, 5}, 5}, []int{1, 2, 3, 4, 5}},
	{rotateArguments{[]int{1, 2, 3, 4, 5}, 0}, []int{1, 2, 3, 4, 5}},
	{rotateArguments{[]int{1, 2, 3, 4, 5}, 3}, []int{3, 4, 5, 1, 2}},
	{rotateArguments{[]int{1, 2, 3, 4, 5}, 8}, []int{3, 4, 5, 1, 2}},
	{rotateArguments{[]int{1, 2, 3, 4, 5}, 4}, []int{2, 3, 4, 5, 1}},
	{rotateArguments{[]int{1, 2, 3, 4, 5}, 1}, []int{5, 1, 2, 3, 4}},
	{rotateArguments{[]int{5}, 10}, []int{5}},
	{rotateArguments{[]int{}, 10}, []int{}},
	{rotateArguments{[]int{1, 2}, 10}, []int{1, 2}},
	{rotateArguments{[]int{1, 2}, 11}, []int{2, 1}},
}

func TestSolutionRotate(t *testing.T) {
	for _, pair := range rotateTests {
		result := SolutionRotate(pair.values.array, pair.values.steps)
		if !IsEqualSlicesOfInt(result, pair.expected) {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", result,
			)
		}
	}
}
