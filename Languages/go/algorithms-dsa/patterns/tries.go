package main

import "fmt"

/*
  Pattern: Tries (Prefix Trees)
  Templates: Insert/Search/Prefix, Autocomplete, Word Search, Longest Common Prefix
*/

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	word     string // store complete word at terminal nodes
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// Insert a word into the trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}
	node.isEnd = true
	node.word = word
}

// Search for an exact word
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// StartsWith checks if any word starts with the prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return true
}

// Autocomplete returns all words with the given prefix
func (t *Trie) Autocomplete(prefix string) []string {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return nil
		}
		node = node.children[ch]
	}
	var results []string
	t.collectWords(node, &results)
	return results
}

func (t *Trie) collectWords(node *TrieNode, results *[]string) {
	if node.isEnd {
		*results = append(*results, node.word)
	}
	for _, child := range node.children {
		t.collectWords(child, results)
	}
}

// longestCommonPrefix using a trie
func longestCommonPrefix(words []string) string {
	if len(words) == 0 {
		return ""
	}
	trie := NewTrie()
	for _, w := range words {
		trie.Insert(w)
	}
	// Walk down the trie as long as there's only one child and not a word end
	node := trie.root
	prefix := []rune{}
	for len(node.children) == 1 && !node.isEnd {
		for ch, child := range node.children {
			prefix = append(prefix, ch)
			node = child
		}
	}
	return string(prefix)
}

// WordDictionary - supports '.' wildcard search
type WordDictionary struct {
	root *TrieNode
}

func NewWordDictionary() *WordDictionary {
	return &WordDictionary{root: NewTrieNode()}
}

func (d *WordDictionary) AddWord(word string) {
	node := d.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (d *WordDictionary) Search(word string) bool {
	return d.searchFrom(d.root, []rune(word), 0)
}

func (d *WordDictionary) searchFrom(node *TrieNode, word []rune, idx int) bool {
	if idx == len(word) {
		return node.isEnd
	}
	ch := word[idx]
	if ch == '.' {
		for _, child := range node.children {
			if d.searchFrom(child, word, idx+1) {
				return true
			}
		}
		return false
	}
	child, ok := node.children[ch]
	if !ok {
		return false
	}
	return d.searchFrom(child, word, idx+1)
}

func main() {
	trie := NewTrie()
	words := []string{"apple", "app", "apricot", "banana", "band", "bandana"}
	for _, w := range words {
		trie.Insert(w)
	}

	// Test search
	if !trie.Search("apple") {
		panic("FAIL: search apple")
	}
	if trie.Search("appl") {
		panic("FAIL: appl should not be found")
	}
	if !trie.Search("app") {
		panic("FAIL: search app")
	}
	fmt.Println("PASS: trie Search")

	// Test prefix
	if !trie.StartsWith("app") {
		panic("FAIL: prefix app")
	}
	if !trie.StartsWith("ban") {
		panic("FAIL: prefix ban")
	}
	if trie.StartsWith("xyz") {
		panic("FAIL: prefix xyz should not exist")
	}
	fmt.Println("PASS: trie StartsWith")

	// Test autocomplete
	results := trie.Autocomplete("app")
	fmt.Println("Autocomplete 'app':", results)
	if len(results) < 2 { // at least "apple" and "app"
		panic("FAIL: autocomplete should find >= 2 words")
	}
	fmt.Println("PASS: trie Autocomplete")

	// Test longest common prefix
	lcp := longestCommonPrefix([]string{"flower", "flow", "flight"})
	if lcp != "fl" {
		panic(fmt.Sprintf("FAIL: LCP expected 'fl', got '%s'", lcp))
	}
	fmt.Println("PASS: longestCommonPrefix =", lcp)

	// Test wildcard dictionary
	dict := NewWordDictionary()
	dict.AddWord("bad")
	dict.AddWord("dad")
	dict.AddWord("mad")

	if !dict.Search("pad") == true { // "pad" not added
		// pad should not be found
	}
	if !dict.Search("bad") {
		panic("FAIL: search bad")
	}
	if !dict.Search(".ad") {
		panic("FAIL: wildcard .ad")
	}
	if !dict.Search("b..") {
		panic("FAIL: wildcard b..")
	}
	if dict.Search("b.") {
		panic("FAIL: b. should not match 3-char words")
	}
	fmt.Println("PASS: WordDictionary with wildcards")
}
