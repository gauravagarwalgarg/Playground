package main

import (
	"fmt"
	"sort"
)

/*
  LC 212 - Word Search II
  Topic: Trie + Backtracking
  Difficulty: Hard
  Time: O(m*n*4^L) | Space: O(sum of word lengths)
*/

type TrieNode struct {
	children [26]*TrieNode
	word     string
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}
	for _, w := range words {
		node := root
		for _, ch := range w {
			idx := ch - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
			}
			node = node.children[idx]
		}
		node.word = w
	}
	m, n := len(board), len(board[0])
	var result []string
	var dfs func(i, j int, node *TrieNode)
	dfs = func(i, j int, node *TrieNode) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == '#' {
			return
		}
		ch := board[i][j]
		next := node.children[ch-'a']
		if next == nil {
			return
		}
		if next.word != "" {
			result = append(result, next.word)
			next.word = "" // avoid duplicates
		}
		board[i][j] = '#'
		dfs(i+1, j, next)
		dfs(i-1, j, next)
		dfs(i, j+1, next)
		dfs(i, j-1, next)
		board[i][j] = ch
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, root)
		}
	}
	sort.Strings(result)
	return result
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	res := findWords(board, []string{"oath", "pea", "eat", "rain"})
	if len(res) == 2 {
		fmt.Println("PASS: findWords")
	} else {
		fmt.Println("FAIL: findWords, got", res)
	}
}
