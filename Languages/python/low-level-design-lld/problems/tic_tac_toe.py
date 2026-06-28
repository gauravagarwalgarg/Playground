"""
Low-Level Design: Tic-Tac-Toe

Board with move validation and O(1) win detection using row/col/diagonal sums.
Player X = +1, Player O = -1. A sum of +3 or -3 means a win.
"""

from dataclasses import dataclass, field


class Board:
    def __init__(self, size: int = 3):
        self.size = size
        self.grid: list[list[int]] = [[0] * size for _ in range(size)]
        self._row_sums = [0] * size
        self._col_sums = [0] * size
        self._diag = 0
        self._anti_diag = 0
        self._moves = 0

    def move(self, row: int, col: int, player: int) -> str:
        """player: +1 (X) or -1 (O). Returns 'X_wins', 'O_wins', 'draw', or 'continue'."""
        if not (0 <= row < self.size and 0 <= col < self.size):
            raise ValueError("Position out of bounds")
        if self.grid[row][col] != 0:
            raise ValueError("Cell already occupied")

        self.grid[row][col] = player
        self._row_sums[row] += player
        self._col_sums[col] += player
        if row == col:
            self._diag += player
        if row + col == self.size - 1:
            self._anti_diag += player
        self._moves += 1

        # Check win
        target = self.size * player
        if (self._row_sums[row] == target or self._col_sums[col] == target
                or self._diag == target or self._anti_diag == target):
            return "X_wins" if player == 1 else "O_wins"

        if self._moves == self.size * self.size:
            return "draw"
        return "continue"

    def display(self) -> str:
        symbols = {0: ".", 1: "X", -1: "O"}
        return "\n".join(" ".join(symbols[c] for c in row) for row in self.grid)


if __name__ == "__main__":
    board = Board()

    # X wins via row
    assert board.move(0, 0, 1) == "continue"   # X
    assert board.move(1, 0, -1) == "continue"  # O
    assert board.move(0, 1, 1) == "continue"   # X
    assert board.move(1, 1, -1) == "continue"  # O
    assert board.move(0, 2, 1) == "X_wins"     # X wins top row

    # O wins via diagonal
    board2 = Board()
    board2.move(0, 0, -1)   # O
    board2.move(0, 1, 1)    # X
    board2.move(1, 1, -1)   # O
    board2.move(0, 2, 1)    # X
    assert board2.move(2, 2, -1) == "O_wins"  # O wins diagonal

    # Draw
    board3 = Board()
    moves = [(0,0,1),(0,1,-1),(0,2,1),(1,0,1),(1,1,-1),(1,2,-1),(2,0,-1),(2,1,1),(2,2,1)]
    results = [board3.move(r, c, p) for r, c, p in moves]
    assert results[-1] == "draw"

    # Invalid move
    try:
        board.move(0, 0, -1)
        assert False, "Should raise"
    except ValueError:
        pass

    print("All tests passed!")
