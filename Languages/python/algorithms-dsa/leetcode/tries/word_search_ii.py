"""
LeetCode #212 - Word Search II
Topic: Tries
Difficulty: Hard

Find all words from dictionary in a grid using Trie + DFS backtracking.

Time Complexity: O(m * n * 4^L) where L = max word length
Space Complexity: O(total characters in words)
"""


class TrieNode:
    def __init__(self) -> None:
        self.children: dict[str, "TrieNode"] = {}
        self.word: str | None = None


def find_words(board: list[list[str]], words: list[str]) -> list[str]:
    root = TrieNode()
    for word in words:
        node = root
        for ch in word:
            if ch not in node.children:
                node.children[ch] = TrieNode()
            node = node.children[ch]
        node.word = word

    rows, cols = len(board), len(board[0])
    result: list[str] = []

    def dfs(r: int, c: int, node: TrieNode) -> None:
        ch = board[r][c]
        if ch not in node.children:
            return
        child = node.children[ch]
        if child.word:
            result.append(child.word)
            child.word = None  # avoid duplicates
        board[r][c] = "#"
        for dr, dc in [(1, 0), (-1, 0), (0, 1), (0, -1)]:
            nr, nc = r + dr, c + dc
            if 0 <= nr < rows and 0 <= nc < cols and board[nr][nc] != "#":
                dfs(nr, nc, child)
        board[r][c] = ch
        if not child.children:
            del node.children[ch]

    for r in range(rows):
        for c in range(cols):
            dfs(r, c, root)
    return result


if __name__ == "__main__":
    board = [["o", "a", "a", "n"], ["e", "t", "a", "e"], ["i", "h", "k", "r"], ["i", "f", "l", "v"]]
    result = find_words(board, ["oath", "pea", "eat", "rain"])
    assert sorted(result) == ["eat", "oath"]

    board2 = [["a", "b"], ["c", "d"]]
    assert find_words(board2, ["abcb"]) == []
    print("All tests passed!")
