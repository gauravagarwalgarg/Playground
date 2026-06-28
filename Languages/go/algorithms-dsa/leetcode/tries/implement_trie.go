/*
LeetCode #208: Implement Trie (Prefix Tree)
Topic: Tries
Difficulty: Medium

Implement a trie with insert, search, and startsWith methods.

Approach: Each node has an array of 26 children (for lowercase letters)
and a boolean marking end of word. Insert builds path, search/startsWith
traverse checking existence.

Time: O(m) per operation where m is word length, Space: O(n * m)
*/
package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func ConstructorTrie() Trie {
	return Trie{root: &TrieNode{}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.findNode(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.findNode(prefix) != nil
}

func (t *Trie) findNode(s string) *TrieNode {
	node := t.root
	for _, ch := range s {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return nil
		}
		node = node.children[idx]
	}
	return node
}

func main() {
	trie := ConstructorTrie()
	trie.Insert("apple")

	if trie.Search("apple") != true {
		fmt.Println("FAIL Test 1")
	} else {
		fmt.Println("PASS Test 1")
	}

	if trie.Search("app") != false {
		fmt.Println("FAIL Test 2")
	} else {
		fmt.Println("PASS Test 2")
	}

	if trie.StartsWith("app") != true {
		fmt.Println("FAIL Test 3")
	} else {
		fmt.Println("PASS Test 3")
	}

	trie.Insert("app")
	if trie.Search("app") != true {
		fmt.Println("FAIL Test 4")
	} else {
		fmt.Println("PASS Test 4")
	}
}
