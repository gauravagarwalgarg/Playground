"""
Low-Level Design: Snake Game

State machine: Snake moves, grows on food, dies on collision.
Board places food randomly and checks boundaries.
"""

from dataclasses import dataclass, field
from collections import deque
import random


@dataclass(frozen=True)
class Position:
    row: int
    col: int


class Snake:
    def __init__(self, start: Position):
        self.body: deque[Position] = deque([start])
        self.direction = (0, 1)  # moving right

    @property
    def head(self) -> Position:
        return self.body[0]

    def set_direction(self, dr: int, dc: int) -> None:
        # Prevent 180-degree reversal
        if (dr + self.direction[0], dc + self.direction[1]) != (0, 0):
            self.direction = (dr, dc)

    def next_head(self) -> Position:
        return Position(self.head.row + self.direction[0], self.head.col + self.direction[1])

    def move(self, grow: bool = False) -> None:
        self.body.appendleft(self.next_head())
        if not grow:
            self.body.pop()

    def collides_self(self) -> bool:
        return self.head in list(self.body)[1:]


class Board:
    def __init__(self, rows: int, cols: int):
        self.rows = rows
        self.cols = cols
        self.food: Position | None = None

    def place_food(self, occupied: set[Position]) -> Position:
        while True:
            pos = Position(random.randint(0, self.rows - 1), random.randint(0, self.cols - 1))
            if pos not in occupied:
                self.food = pos
                return pos

    def is_out_of_bounds(self, pos: Position) -> bool:
        return not (0 <= pos.row < self.rows and 0 <= pos.col < self.cols)


class Game:
    def __init__(self, rows: int = 10, cols: int = 10):
        self.board = Board(rows, cols)
        self.snake = Snake(Position(rows // 2, cols // 2))
        self.score = 0
        self.game_over = False
        self.board.place_food(set(self.snake.body))

    def tick(self) -> str:
        """Advance one step. Returns status."""
        if self.game_over:
            return "game_over"
        next_pos = self.snake.next_head()
        if self.board.is_out_of_bounds(next_pos):
            self.game_over = True
            return "wall_collision"
        grow = next_pos == self.board.food
        self.snake.move(grow=grow)
        if self.snake.collides_self():
            self.game_over = True
            return "self_collision"
        if grow:
            self.score += 1
            self.board.place_food(set(self.snake.body))
            return "ate_food"
        return "moved"


if __name__ == "__main__":
    random.seed(42)
    game = Game(5, 5)
    start_pos = game.snake.head

    # Move right
    status = game.tick()
    assert status == "moved" or status == "ate_food"
    assert game.snake.head.col == start_pos.col + 1

    # Change direction down
    game.snake.set_direction(1, 0)
    game.tick()
    assert game.snake.head.row == start_pos.row + 1

    # Can't reverse direction (go up when going down)
    game.snake.set_direction(-1, 0)
    assert game.snake.direction == (1, 0)  # unchanged

    # Wall collision test
    small_game = Game(3, 3)
    small_game.snake = Snake(Position(1, 2))  # near right wall
    small_game.snake.set_direction(0, 1)
    status = small_game.tick()
    assert status == "wall_collision"
    assert small_game.game_over

    print("All tests passed!")
