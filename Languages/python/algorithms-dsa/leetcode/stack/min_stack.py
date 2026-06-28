"""
LeetCode #155 - Min Stack
Topic: Stack
Difficulty: Medium

Design a stack that supports push, pop, top, and getMin in O(1).

Time Complexity: O(1) for all operations
Space Complexity: O(n)
"""


class MinStack:
    def __init__(self) -> None:
        self.stack: list[tuple[int, int]] = []

    def push(self, val: int) -> None:
        current_min = min(val, self.stack[-1][1]) if self.stack else val
        self.stack.append((val, current_min))

    def pop(self) -> None:
        self.stack.pop()

    def top(self) -> int:
        return self.stack[-1][0]

    def get_min(self) -> int:
        return self.stack[-1][1]


if __name__ == "__main__":
    ms = MinStack()
    ms.push(-2)
    ms.push(0)
    ms.push(-3)
    assert ms.get_min() == -3
    ms.pop()
    assert ms.top() == 0
    assert ms.get_min() == -2
    print("All tests passed!")
