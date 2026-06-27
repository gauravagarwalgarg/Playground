/*
 * LeetCode 424 - Longest Repeating Character Replacement
 * Topic: Sliding Window
 * Difficulty: Medium
 *
 * Sliding window: track frequency of each char. Window is valid if
 * (window_size - maxFreq) <= k.
 * Time: O(n)
 * Space: O(1) (26 letters)
 */
#include <iostream>
#include <string>
#include <vector>
#include <cassert>
using namespace std;

int characterReplacement(string s, int k) {
    vector<int> count(26, 0);
    int left = 0, maxFreq = 0, result = 0;
    for (int right = 0; right < (int)s.size(); right++) {
        count[s[right] - 'A']++;
        maxFreq = max(maxFreq, count[s[right] - 'A']);
        while ((right - left + 1) - maxFreq > k) {
            count[s[left] - 'A']--;
            left++;
        }
        result = max(result, right - left + 1);
    }
    return result;
}

int main() {
    assert(characterReplacement("ABAB", 2) == 4);
    assert(characterReplacement("AABABBA", 1) == 4);
    assert(characterReplacement("AAAA", 0) == 4);
    assert(characterReplacement("ABCD", 2) == 3);

    cout << "All tests passed!" << endl;
    return 0;
}
