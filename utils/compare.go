package utils

func CompareStringSlices(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}

	for i, str := range left {
		if right[i] != str {
			return false
		}
	}

	return true
}
