/*
 * LeetCode 49 - Group Anagrams
 * Topic: Arrays & Hashing
 * Difficulty: Medium
 * 
 * Sort each word to use as a key, group words with same sorted key.
 * Time: O(n * k * log k) where n = number of strings, k = max string length
 * Space: O(n * k)
 */
#include <iostream>
#include <vector>
#include <string>
#include <unordered_map>
#include <algorithm>
#include <cassert>
using namespace std;

vector<vector<string>> groupAnagrams(vector<string>& strs) {
    unordered_map<string, vector<string>> mp;
    for (auto& s : strs) {
        string key = s;
        sort(key.begin(), key.end());
        mp[key].push_back(s);
    }
    vector<vector<string>> res;
    for (auto& [k, v] : mp) res.push_back(v);
    return res;
}

int main() {
    vector<string> t1 = {"eat","tea","tan","ate","nat","bat"};
    auto r1 = groupAnagrams(t1);
    assert(r1.size() == 3);

    vector<string> t2 = {""};
    auto r2 = groupAnagrams(t2);
    assert(r2.size() == 1 && r2[0][0] == "");

    vector<string> t3 = {"a"};
    auto r3 = groupAnagrams(t3);
    assert(r3.size() == 1 && r3[0][0] == "a");

    vector<string> t4 = {"abc","bca","cab","xyz","zyx"};
    auto r4 = groupAnagrams(t4);
    assert(r4.size() == 2);

    cout << "All tests passed!" << endl;
    return 0;
}
