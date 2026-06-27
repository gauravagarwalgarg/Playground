"""
LeetCode #79 - Word Search
Topic: Backtracking
Difficulty: Medium

Search for a word in a grid using DFS backtracking.

Time Complexity: O(m * n * 4^L) where L = word length
Space Complexity: O(L) recursion depth
"""


def exist(board: list[list[str]], word: str) -> bool:
    rows, cols = len(board), len(board[0])

    def dfs(r: int, c: int, i: int) -> bool:
        if i == len(word):
            return True
        if r < 0 or r >= rows or c < 0 or c >= cols or board[r][c] != word[i]:
            return False
        temp = board[r][c]
        board[r][c] = "#"
        found = (
            dfs(r + 1, c, i + 1) or dfs(r - 1, c, i + 1) or
            dfs(r, c + 1, i + 1) or dfs(r, c - 1, i + 1)
        )
        board[r][c] = temp
        return found

    for r in range(rows):
        for c in range(cols):
            if dfs(r, c, 0):
                return True
    return False


if __name__ == "__main__":
    board = [["A", "B", "C", "E"], ["S", "F", "C", "S"], ["A", "D", "E", "E"]]
    assert exist(board, "ABCCED") is True
    assert exist(board, "SEE") is True
    assert exist(board, "ABCB") is False
    print("All tests passed!")
