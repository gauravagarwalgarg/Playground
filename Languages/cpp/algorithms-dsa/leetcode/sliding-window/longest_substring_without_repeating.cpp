/**
 * LeetCode 3: Longest Substring Without Repeating Characters
 * Topic: Sliding Window
 * Difficulty: Medium
 *
 * Given a string s, find the length of the longest substring without repeating characters.
 *
 * Approach: Sliding window with a hashset. Expand the right pointer, if a duplicate is
 * found shrink from the left until the duplicate is removed. Track the max window size.
 *
 * Time: O(n), Space: O(min(n, alphabet))
 */
#include <iostream>
#include <string>
#include <unordered_set>
#include <algorithm>
#include <cassert>
using namespace std;

int lengthOfLongestSubstring(string s) {
    unordered_set<char> charSet;
    int left = 0;
    int maxLen = 0;

    for (int right = 0; right < (int)s.size(); right++) {
        while (charSet.count(s[right])) {
            charSet.erase(s[left]);
            left++;
        }
        charSet.insert(s[right]);
        maxLen = max(maxLen, right - left + 1);
    }

    return maxLen;
}

int main() {
    assert(lengthOfLongestSubstring("abcabcbb") == 3);
    assert(lengthOfLongestSubstring("bbbbb") == 1);
    assert(lengthOfLongestSubstring("pwwkew") == 3);
    assert(lengthOfLongestSubstring("") == 0);
    assert(lengthOfLongestSubstring(" ") == 1);
    assert(lengthOfLongestSubstring("au") == 2);
    assert(lengthOfLongestSubstring("dvdf") == 3);

    cout << "All tests passed!" << endl;
    return 0;
}
