"""
LeetCode #125 - Valid Palindrome
Topic: Two Pointers
Difficulty: Easy

Check if a string is a palindrome considering only alphanumeric characters.

Time Complexity: O(n)
Space Complexity: O(1)
"""


def is_palindrome(s: str) -> bool:
    left, right = 0, len(s) - 1
    while left < right:
        while left < right and not s[left].isalnum():
            left += 1
        while left < right and not s[right].isalnum():
            right -= 1
        if s[left].lower() != s[right].lower():
            return False
        left += 1
        right -= 1
    return True


if __name__ == "__main__":
    assert is_palindrome("A man, a plan, a canal: Panama") is True
    assert is_palindrome("race a car") is False
    assert is_palindrome(" ") is True
    assert is_palindrome("0P") is False
    print("All tests passed!")
