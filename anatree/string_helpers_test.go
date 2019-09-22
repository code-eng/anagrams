package anatree

import "testing"

var toFrequenciesTests = []struct {
	input    string
	expected []letterFrequency
}{
	{
		"",
		[]letterFrequency{},
	},
	{
		"1@#-+  ",
		[]letterFrequency{},
	},
	{
		"aaBB",
		[]letterFrequency{{'a', 2}, {'b', 2}},
	},
	{
		"!21AA   BB(*(_#",
		[]letterFrequency{{'a', 2}, {'b', 2}},
	},
	{
		"sasha", // aahss - sorted sasha
		[]letterFrequency{{'a', 2}, {'h', 1}, {'s', 2}},
	},
	{
		"S  a S H !123 a #_(  .)(.  )_#", // aahss - sorted sasha
		[]letterFrequency{{'a', 2}, {'h', 1}, {'s', 2}},
	},
}

func compareFrequencies(left, right []letterFrequency) bool {
	if len(left) != len(right) {
		return false
	}

	for i, frequency := range left {
		if right[i] != frequency {
			return false
		}
	}

	return true
}

func TestToFrequencies(t *testing.T) {
	for _, testCase := range toFrequenciesTests {
		result := toFrequencies(testCase.input)
		if !compareFrequencies(testCase.expected, result) {
			t.Errorf("\nExpected: %v\nGot: %v\n", testCase.expected, result)
		}
	}
}
