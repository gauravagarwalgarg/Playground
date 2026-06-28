"""
LeetCode #51 - N-Queens
Topic: Backtracking
Difficulty: Hard

Place N queens on an N×N board row-by-row with column/diagonal tracking.

Time Complexity: O(n!)
Space Complexity: O(n)
"""


def solve_n_queens(n: int) -> list[list[str]]:
    result: list[list[str]] = []
    cols: set[int] = set()
    diag1: set[int] = set()  # row - col
    diag2: set[int] = set()  # row + col
    board = [["." for _ in range(n)] for _ in range(n)]

    def backtrack(row: int) -> None:
        if row == n:
            result.append(["".join(r) for r in board])
            return
        for col in range(n):
            if col in cols or (row - col) in diag1 or (row + col) in diag2:
                continue
            cols.add(col)
            diag1.add(row - col)
            diag2.add(row + col)
            board[row][col] = "Q"
            backtrack(row + 1)
            board[row][col] = "."
            cols.remove(col)
            diag1.remove(row - col)
            diag2.remove(row + col)

    backtrack(0)
    return result


if __name__ == "__main__":
    assert len(solve_n_queens(4)) == 2
    assert len(solve_n_queens(1)) == 1
    assert solve_n_queens(1) == [["Q"]]
    print("All tests passed!")
