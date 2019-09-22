package anatree

import (
	"anagrams/utils"
	"testing"
)

var sampleDict = []string{"foobar", "aabb", "baba", "boofar", "test"}

var anatreeTests = []struct {
	dictionary []string
	input      string
	expected   []string
}{
	{
		[]string{},
		"any",
		[]string{},
	},
	{
		[]string{"123", "#*&$", "   "},
		"any",
		[]string{},
	},
	{
		sampleDict,
		"any",
		[]string{},
	},
	{
		sampleDict,
		"!1$52315)  -",
		[]string{},
	},
	{
		sampleDict,
		"",
		[]string{},
	},
	{
		sampleDict,
		"foobar",
		[]string{"foobar", "boofar"},
	},
	{
		sampleDict,
		"f  $oObaR!",
		[]string{"foobar", "boofar"},
	},
	{
		sampleDict,
		"test",
		[]string{"test"},
	},
}

func TestAnatree(t *testing.T) {
	// This test covers all three methods of Anatree:
	// FromWords
	// AddWord
	// GetAnagrams

	for _, testCase := range anatreeTests {
		tree := FromWords(testCase.dictionary)
		result := tree.GetAnagrams(testCase.input)
		if !utils.CompareStringSlices(testCase.expected, result) {
			t.Errorf("\nExpected: %v\n Got: %v\n", testCase.expected, result)
		}
	}
}
