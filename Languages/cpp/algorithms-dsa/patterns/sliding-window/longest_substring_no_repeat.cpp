/**
 * Pattern: Sliding Window
 * Problem: Longest Substring Without Repeating Characters (LeetCode 3)
 *
 * Time: O(n), Space: O(min(n, alphabet_size))
 */
#include <iostream>
#include <string>
#include <unordered_set>
#include <algorithm>
#include <cassert>
using namespace std;

int lengthOfLongestSubstring(string s) {
    unordered_set<char> window;
    int left = 0, maxLen = 0;
    for (int right = 0; right < (int)s.size(); right++) {
        while (window.count(s[right])) {
            window.erase(s[left]);
            left++;
        }
        window.insert(s[right]);
        maxLen = max(maxLen, right - left + 1);
    }
    return maxLen;
}

int main() {
    assert(lengthOfLongestSubstring("abcabcbb") == 3);
    assert(lengthOfLongestSubstring("bbbbb") == 1);
    assert(lengthOfLongestSubstring("pwwkew") == 3);
    assert(lengthOfLongestSubstring("") == 0);

    cout << "All tests passed!" << endl;
    return 0;
}
