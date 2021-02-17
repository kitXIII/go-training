package konovalov

// IsEqualSlicesOfInt shallow compare 2 slices of integer arrays and return true or false
func IsEqualSlicesOfInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
