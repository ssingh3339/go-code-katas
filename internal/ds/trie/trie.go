package trie

// TrieNode represents a node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie represents a prefix tree
type Trie struct {
	root *TrieNode
}

// NewTrie creates and returns a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// Insert adds a word to the Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search checks if a word exists in the Trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// StartsWith checks if there is any word in the Trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return true
}

// Delete removes a word from the Trie
func (t *Trie) Delete(word string) bool {
	return t.deleteHelper(t.root, word, 0)
}

func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		if !node.isEnd {
			return false
		}
		node.isEnd = false
		return len(node.children) == 0
	}

	ch := rune(word[index])
	childNode, exists := node.children[ch]
	if !exists {
		return false
	}

	shouldDeleteChild := t.deleteHelper(childNode, word, index+1)

	if shouldDeleteChild {
		delete(node.children, ch)
		return len(node.children) == 0 && !node.isEnd
	}

	return false
}
