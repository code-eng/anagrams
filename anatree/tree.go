package anatree

type Anatree map[letterFrequency]*node

type letterFrequency struct {
	letter    rune
	frequency int
}

type node struct {
	words []string
	tree  Anatree
}

func createNode() *node {
	return &node{make([]string, 0), Anatree{}}
}

func FromWords(words []string) Anatree {
	tree := Anatree{}

	for _, word := range words {
		tree.AddWord(word)
	}

	return tree
}

func (tree Anatree) GetAnagrams(word string) []string {
	charsFrequency := toFrequencies(word)

	if len(charsFrequency) == 0 {
		return []string{}
	}

	currentNode := tree[charsFrequency[0]]
	for _, frequency := range charsFrequency[1:] {
		if currentNode == nil {
			return []string{}
		}
		currentNode = currentNode.tree[frequency]
	}

	return currentNode.words
}

func (tree Anatree) AddWord(word string) {
	frequencies := toFrequencies(word)
	if len(frequencies) == 0 {
		return
	}

	if tree[frequencies[0]] == nil {
		tree[frequencies[0]] = createNode()
	}

	currentNode := tree[frequencies[0]]

	for _, frequency := range frequencies[1:] {

		if currentNode.tree[frequency] == nil {
			currentNode.tree[frequency] = createNode()
		}

		currentNode = currentNode.tree[frequency]
	}

	currentNode.words = append(currentNode.words, word)
}
