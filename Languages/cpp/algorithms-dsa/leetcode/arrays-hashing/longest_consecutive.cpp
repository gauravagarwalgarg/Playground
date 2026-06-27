/*
 * LeetCode 128 - Longest Consecutive Sequence
 * Topic: Arrays & Hashing
 * Difficulty: Medium
 *
 * Insert all numbers into a HashSet. For each number that is a sequence start
 * (no num-1 in set), count consecutive elements.
 * Time: O(n)
 * Space: O(n)
 */
#include <iostream>
#include <vector>
#include <unordered_set>
#include <cassert>
using namespace std;

int longestConsecutive(vector<int>& nums) {
    unordered_set<int> s(nums.begin(), nums.end());
    int longest = 0;
    for (int n : s) {
        if (s.find(n - 1) == s.end()) {
            int len = 1;
            while (s.count(n + len)) len++;
            longest = max(longest, len);
        }
    }
    return longest;
}

int main() {
    vector<int> t1 = {100,4,200,1,3,2};
    assert(longestConsecutive(t1) == 4);

    vector<int> t2 = {0,3,7,2,5,8,4,6,0,1};
    assert(longestConsecutive(t2) == 9);

    vector<int> t3 = {};
    assert(longestConsecutive(t3) == 0);

    vector<int> t4 = {1,2,0,1};
    assert(longestConsecutive(t4) == 3);

    cout << "All tests passed!" << endl;
    return 0;
}
