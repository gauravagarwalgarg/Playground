/*
 * LeetCode 76 - Minimum Window Substring
 * Topic: Sliding Window
 * Difficulty: Hard
 *
 * Expand window to include all chars of t, then contract to find minimum.
 * Use frequency maps and a "formed" counter.
 * Time: O(n)
 * Space: O(n)
 */
#include <iostream>
#include <string>
#include <unordered_map>
#include <climits>
#include <cassert>
using namespace std;

string minWindow(string s, string t) {
    unordered_map<char, int> need, have;
    for (char c : t) need[c]++;
    int required = need.size(), formed = 0;
    int left = 0, minLen = INT_MAX, minStart = 0;

    for (int right = 0; right < (int)s.size(); right++) {
        have[s[right]]++;
        if (need.count(s[right]) && have[s[right]] == need[s[right]])
            formed++;
        while (formed == required) {
            int len = right - left + 1;
            if (len < minLen) { minLen = len; minStart = left; }
            have[s[left]]--;
            if (need.count(s[left]) && have[s[left]] < need[s[left]])
                formed--;
            left++;
        }
    }
    return minLen == INT_MAX ? "" : s.substr(minStart, minLen);
}

int main() {
    assert(minWindow("ADOBECODEBANC", "ABC") == "BANC");
    assert(minWindow("a", "a") == "a");
    assert(minWindow("a", "aa") == "");
    assert(minWindow("aa", "aa") == "aa");

    cout << "All tests passed!" << endl;
    return 0;
}
