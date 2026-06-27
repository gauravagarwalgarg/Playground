/*
 * LeetCode 125 - Valid Palindrome
 * Topic: Two Pointers
 * Difficulty: Easy
 *
 * Two pointers from both ends, skip non-alphanumeric, compare lowercase.
 * Time: O(n)
 * Space: O(1)
 */
#include <iostream>
#include <string>
#include <cassert>
using namespace std;

bool isPalindrome(string s) {
    int left = 0, right = s.size() - 1;
    while (left < right) {
        while (left < right && !isalnum(s[left])) left++;
        while (left < right && !isalnum(s[right])) right--;
        if (tolower(s[left]) != tolower(s[right])) return false;
        left++; right--;
    }
    return true;
}

int main() {
    assert(isPalindrome("A man, a plan, a canal: Panama") == true);
    assert(isPalindrome("race a car") == false);
    assert(isPalindrome(" ") == true);
    assert(isPalindrome("0P") == false);

    cout << "All tests passed!" << endl;
    return 0;
}
