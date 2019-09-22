package state

import (
	"anagrams/anatree"
	"sync"
)

type dictionary struct {
	mu   sync.Mutex
	tree anatree.Anatree
}

var dict = dictionary{
	mu:   sync.Mutex{},
	tree: anatree.Anatree{},
}

func LoadDictionary(words []string) {
	tree := anatree.FromWords(words)

	dict.mu.Lock()
	defer dict.mu.Unlock()
	dict.tree = tree
}

func GetAnagrams(word string) []string {
	return dict.tree.GetAnagrams(word)
}
