"""
LeetCode 70: Climbing Stairs
You are climbing a staircase. It takes n steps to reach the top.
Each time you can climb 1 or 2 steps. How many distinct ways can you climb?

Time: O(n), Space: O(1)
"""


def climbing_stairs(n: int) -> int:
    if n <= 2:
        return n
    prev, curr = 1, 2
    for _ in range(3, n + 1):
        prev, curr = curr, prev + curr
    return curr


if __name__ == "__main__":
    assert climbing_stairs(1) == 1
    assert climbing_stairs(2) == 2
    assert climbing_stairs(3) == 3
    assert climbing_stairs(4) == 5
    assert climbing_stairs(5) == 8
    assert climbing_stairs(10) == 89

    print("All tests passed!")
