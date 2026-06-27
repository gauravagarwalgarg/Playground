"""
Trie (Prefix Tree) Pattern Template

Complete Trie implementation with:
- insert(word): add word to trie
- search(word): check if exact word exists
- starts_with(prefix): check if any word starts with prefix
- Bonus: autocomplete, word search with wildcards

Use when: "prefix matching", "autocomplete", "word dictionary", "spell checker",
          "word search board", "longest common prefix", "counting prefixes"

Run: python trie.py
"""

from typing import List, Optional


class TrieNode:
    """A node in the Trie. Each node has up to 26 children (a-z)."""

    def __init__(self):
        self.children = {}      # char → TrieNode
        self.is_end = False     # marks end of a complete word
        self.word_count = 0     # number of words ending here (for duplicates)
        self.prefix_count = 0   # number of words passing through this node


class Trie:
    """
    Prefix tree for efficient string operations.
    
    Time Complexity (where L = word length):
    - insert: O(L)
    - search: O(L)
    - starts_with: O(L)
    
    Space Complexity: O(total characters across all words)
    """

    def __init__(self):
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        """Insert a word into the trie."""
        node = self.root

        for char in word:
            if char not in node.children:
                node.children[char] = TrieNode()
            node = node.children[char]
            node.prefix_count += 1  # this node is part of a word's prefix

        node.is_end = True
        node.word_count += 1

    def search(self, word: str) -> bool:
        """Return True if word is in the trie (exact match)."""
        node = self._find_node(word)
        return node is not None and node.is_end

    def starts_with(self, prefix: str) -> bool:
        """Return True if any word in trie starts with the given prefix."""
        return self._find_node(prefix) is not None

    def count_words_with_prefix(self, prefix: str) -> int:
        """Count how many words have the given prefix."""
        node = self._find_node(prefix)
        return node.prefix_count if node else 0

    def _find_node(self, prefix: str) -> Optional[TrieNode]:
        """Navigate trie to end of prefix. Returns None if prefix doesn't exist."""
        node = self.root
        for char in prefix:
            if char not in node.children:
                return None
            node = node.children[char]
        return node

    # =========================================================================
    # BONUS: Autocomplete find all words with given prefix
    # =========================================================================
    def autocomplete(self, prefix: str) -> List[str]:
        """Return all words in trie that start with prefix."""
        node = self._find_node(prefix)
        if node is None:
            return []

        results = []
        self._collect_words(node, list(prefix), results)
        return results

    def _collect_words(self, node: TrieNode, path: List[str], results: List[str]):
        """DFS to collect all words from this node."""
        if node.is_end:
            results.append(''.join(path))

        for char, child in sorted(node.children.items()):
            path.append(char)
            self._collect_words(child, path, results)
            path.pop()

    # =========================================================================
    # BONUS: Search with wildcards (. matches any character)
    # Used in: LeetCode "Add and Search Word" / Word Dictionary
    # =========================================================================
    def search_wildcard(self, word: str) -> bool:
        """Search with '.' as wildcard matching any single character."""

        def dfs(node: TrieNode, i: int) -> bool:
            if i == len(word):
                return node.is_end

            char = word[i]
            if char == '.':
                # Try all children
                for child in node.children.values():
                    if dfs(child, i + 1):
                        return True
                return False
            else:
                if char not in node.children:
                    return False
                return dfs(node.children[char], i + 1)

        return dfs(self.root, 0)

    # =========================================================================
    # BONUS: Delete a word from trie
    # =========================================================================
    def delete(self, word: str) -> bool:
        """Delete word from trie. Returns True if word existed."""

        def _delete(node: TrieNode, word: str, depth: int) -> bool:
            if depth == len(word):
                if not node.is_end:
                    return False
                node.is_end = False
                node.word_count -= 1
                return True  # can delete this node if no children

            char = word[depth]
            if char not in node.children:
                return False

            child = node.children[char]
            should_delete = _delete(child, word, depth + 1)

            if should_delete:
                child.prefix_count -= 1
                # Remove child if it has no words passing through
                if child.prefix_count == 0:
                    del node.children[char]

            return should_delete

        return _delete(self.root, word, 0)


# =============================================================================
# APPLICATION: Longest Common Prefix using Trie
# =============================================================================
def longest_common_prefix(words: List[str]) -> str:
    """Find longest common prefix among all words."""
    if not words:
        return ""

    trie = Trie()
    for word in words:
        trie.insert(word)

    # Follow the path as long as there's exactly one child and not end of word
    prefix = []
    node = trie.root

    while len(node.children) == 1 and not node.is_end:
        char = next(iter(node.children))
        prefix.append(char)
        node = node.children[char]

    return ''.join(prefix)


# =============================================================================
if __name__ == "__main__":
    # Basic operations
    trie = Trie()

    words = ["apple", "app", "apricot", "banana", "band", "bandana"]
    for word in words:
        trie.insert(word)

    print("=== Basic Operations ===")
    print(f"search('apple'): {trie.search('apple')}")         # True
    print(f"search('app'): {trie.search('app')}")             # True
    print(f"search('ap'): {trie.search('ap')}")               # False (not complete word)
    print(f"starts_with('ap'): {trie.starts_with('ap')}")     # True
    print(f"starts_with('xyz'): {trie.starts_with('xyz')}")   # False

    # Count words with prefix
    print(f"\n=== Prefix Counting ===")
    print(f"Words with prefix 'app': {trie.count_words_with_prefix('app')}")    # 2
    print(f"Words with prefix 'ban': {trie.count_words_with_prefix('ban')}")    # 3

    # Autocomplete
    print(f"\n=== Autocomplete ===")
    print(f"autocomplete('ap'): {trie.autocomplete('ap')}")
    print(f"autocomplete('ban'): {trie.autocomplete('ban')}")

    # Wildcard search
    print(f"\n=== Wildcard Search ===")
    print(f"search_wildcard('a.ple'): {trie.search_wildcard('a.ple')}")   # True (apple)
    print(f"search_wildcard('b..d'): {trie.search_wildcard('b..d')}")     # True (band)
    print(f"search_wildcard('b....'): {trie.search_wildcard('b....')}")   # True (banda? no, banana)

    # Delete
    print(f"\n=== Delete ===")
    print(f"delete('app'): {trie.delete('app')}")
    print(f"search('app'): {trie.search('app')}")         # False
    print(f"search('apple'): {trie.search('apple')}")     # True (still exists)

    # Longest common prefix
    print(f"\n=== Longest Common Prefix ===")
    print(f"LCP ['flower','flow','flight']: '{longest_common_prefix(['flower', 'flow', 'flight'])}'")
