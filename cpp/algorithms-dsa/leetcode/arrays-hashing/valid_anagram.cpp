/**
 * LeetCode 242: Valid Anagram
 * Topic: Arrays & Hashing
 * Difficulty: Easy
 *
 * Return true if t is an anagram of s.
 * Time: O(n), Space: O(1) -- fixed 26-char alphabet
 */
#include <iostream>
#include <string>
#include <array>
#include <cassert>
using namespace std;

bool isAnagram(string s, string t) {
    if (s.size() != t.size()) return false;
    array<int, 26> count{};
    for (int i = 0; i < (int)s.size(); i++) {
        count[s[i] - 'a']++;
        count[t[i] - 'a']--;
    }
    for (int c : count) {
        if (c != 0) return false;
    }
    return true;
}

int main() {
    assert(isAnagram("anagram", "nagaram") == true);
    assert(isAnagram("rat", "car") == false);
    assert(isAnagram("a", "a") == true);

    cout << "All tests passed!" << endl;
    return 0;
}
