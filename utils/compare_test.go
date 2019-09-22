package utils

import "testing"

func TestCompareStringSlices(t *testing.T) {
	compareStringSlicesTests := []struct {
		left     []string
		right    []string
		expected bool
	}{
		{
			[]string{"a", "b"},
			[]string{},
			false,
		},
		{
			[]string{},
			[]string{"a", "b"},
			false,
		},
		{
			[]string{"b", "a"},
			[]string{"a", "b"},
			false,
		},
		{
			[]string{"a", "b"},
			[]string{"a", "b"},
			true,
		},
	}

	for _, testCase := range compareStringSlicesTests {
		res := CompareStringSlices(testCase.left, testCase.right)
		if testCase.expected != res {
			t.Errorf(
				"\nExpected: %v\nGot: %v\nTestcase: %v",
				testCase.expected,
				res,
				testCase,
			)
		}
	}
}
