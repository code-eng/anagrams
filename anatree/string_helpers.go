package anatree

import (
	"sort"
	"unicode"
)

func stringToLetters(s string) []rune {
	res := make([]rune, len(s))
	for _, runeValue := range s {
		if unicode.IsLetter(runeValue) {
			res = append(res, unicode.ToLower(runeValue))
		}
	}
	return res
}

func sortRunes(runes []rune) []rune {
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return runes
}

func toFrequencies(word string) []letterFrequency {
	chars := stringToLetters(word)
	sortRunes(chars)
	// we want to preserve order, that's why we are not using map
	return buildFrequencies(chars)
}

func buildFrequencies(runes []rune) (res []letterFrequency) {
	var currentChar rune
	var count int
	for _, char := range runes {
		switch {
		case currentChar == 0:
			currentChar = char
			count = 1
		case currentChar == char:
			count += 1
		case currentChar != char:
			res = append(res, letterFrequency{currentChar, count})
			count = 1
			currentChar = char
		}
	}

	if currentChar != 0 {
		res = append(res, letterFrequency{currentChar, count})
	}

	return
}
