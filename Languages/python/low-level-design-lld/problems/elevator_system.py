"""
Low-Level Design: Elevator System

Requirements:
- Handle multiple requests with floor and direction
- Move elevator toward requests using a scan (LOOK) algorithm
- Track current floor, direction, and pending requests

Classes: Elevator, Request, ElevatorController
"""

from enum import Enum
from dataclasses import dataclass
import heapq


class Direction(Enum):
    UP = 1
    DOWN = -1
    IDLE = 0


@dataclass(order=True)
class Request:
    floor: int
    direction: Direction


class Elevator:
    def __init__(self, elevator_id: int, min_floor: int = 0, max_floor: int = 10):
        self.id = elevator_id
        self.current_floor = 0
        self.direction = Direction.IDLE
        self.min_floor = min_floor
        self.max_floor = max_floor

    def move_to(self, floor: int) -> list[int]:
        """Returns list of floors visited while moving to target."""
        visited = []
        step = 1 if floor > self.current_floor else -1
        while self.current_floor != floor:
            self.current_floor += step
            visited.append(self.current_floor)
        return visited


class ElevatorController:
    def __init__(self, elevator: Elevator):
        self.elevator = elevator
        self._up_stops: list[int] = []  # min-heap
        self._down_stops: list[int] = []  # max-heap (negated)

    def add_request(self, request: Request) -> None:
        if request.direction == Direction.UP:
            heapq.heappush(self._up_stops, request.floor)
        else:
            heapq.heappush(self._down_stops, -request.floor)

    def process_next(self) -> int | None:
        """Process next stop based on current direction (LOOK algorithm)."""
        if self.elevator.direction in (Direction.UP, Direction.IDLE):
            if self._up_stops:
                target = heapq.heappop(self._up_stops)
                self.elevator.direction = Direction.UP
                self.elevator.move_to(target)
                if not self._up_stops:
                    self.elevator.direction = Direction.DOWN if self._down_stops else Direction.IDLE
                return target
        if self.elevator.direction in (Direction.DOWN, Direction.IDLE):
            if self._down_stops:
                target = -heapq.heappop(self._down_stops)
                self.elevator.direction = Direction.DOWN
                self.elevator.move_to(target)
                if not self._down_stops:
                    self.elevator.direction = Direction.UP if self._up_stops else Direction.IDLE
                return target
        self.elevator.direction = Direction.IDLE
        return None

    @property
    def pending(self) -> int:
        return len(self._up_stops) + len(self._down_stops)


if __name__ == "__main__":
    elevator = Elevator(1, min_floor=0, max_floor=10)
    controller = ElevatorController(elevator)

    controller.add_request(Request(5, Direction.UP))
    controller.add_request(Request(3, Direction.UP))
    controller.add_request(Request(7, Direction.DOWN))

    # Processes UP requests in order (nearest first)
    assert controller.process_next() == 3
    assert elevator.current_floor == 3

    assert controller.process_next() == 5
    assert elevator.current_floor == 5

    # Switches to DOWN
    assert controller.process_next() == 7
    assert elevator.current_floor == 7
    assert elevator.direction == Direction.IDLE

    assert controller.process_next() is None
    assert controller.pending == 0

    print("All tests passed!")
