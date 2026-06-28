/**
 * LeetCode 424: Longest Repeating Character Replacement
 * Topic: Sliding Window
 * Difficulty: Medium
 *
 * You are given a string s and an integer k. You can choose any character of the string
 * and change it to any other uppercase English character. You can perform this operation
 * at most k times. Return the length of the longest substring containing the same letter
 * you can get after performing the above operations.
 *
 * Approach: Sliding window tracking the frequency of each character in the window.
 * The key insight is: window_size - max_frequency <= k means the window is valid
 * (we can replace the non-max characters with at most k operations).
 * Shrink from left when window becomes invalid.
 *
 * Time: O(n), Space: O(26) = O(1)
 */
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
#include <cassert>
using namespace std;

int characterReplacement(string s, int k) {
    vector<int> count(26, 0);
    int maxFreq = 0;
    int left = 0;
    int maxLen = 0;

    for (int right = 0; right < (int)s.size(); right++) {
        count[s[right] - 'A']++;
        maxFreq = max(maxFreq, count[s[right] - 'A']);

        // If window is invalid, shrink from left
        while ((right - left + 1) - maxFreq > k) {
            count[s[left] - 'A']--;
            left++;
        }

        maxLen = max(maxLen, right - left + 1);
    }

    return maxLen;
}

int main() {
    assert(characterReplacement("ABAB", 2) == 4);
    assert(characterReplacement("AABABBA", 1) == 4);
    assert(characterReplacement("AAAA", 2) == 4);
    assert(characterReplacement("ABCD", 0) == 1);
    assert(characterReplacement("ABCD", 1) == 2);
    assert(characterReplacement("A", 0) == 1);
    assert(characterReplacement("", 2) == 0);

    cout << "All tests passed!" << endl;
    return 0;
}
