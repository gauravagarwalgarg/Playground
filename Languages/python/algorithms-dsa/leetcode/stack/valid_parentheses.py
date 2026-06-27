"""
LeetCode #20 - Valid Parentheses
Topic: Stack
Difficulty: Easy

Check if brackets are balanced using stack matching.

Time Complexity: O(n)
Space Complexity: O(n)
"""


def is_valid(s: str) -> bool:
    stack: list[str] = []
    pairs = {")": "(", "}": "{", "]": "["}
    for ch in s:
        if ch in pairs:
            if not stack or stack[-1] != pairs[ch]:
                return False
            stack.pop()
        else:
            stack.append(ch)
    return len(stack) == 0


if __name__ == "__main__":
    assert is_valid("()") is True
    assert is_valid("()[]{}") is True
    assert is_valid("(]") is False
    assert is_valid("([)]") is False
    assert is_valid("{[]}") is True
    print("All tests passed!")
