/*
 * LeetCode 347 - Top K Frequent Elements
 * Topic: Arrays & Hashing
 * Difficulty: Medium
 *
 * Use bucket sort: index = frequency, value = list of elements with that frequency.
 * Time: O(n)
 * Space: O(n)
 */
#include <iostream>
#include <vector>
#include <unordered_map>
#include <cassert>
#include <algorithm>
using namespace std;

vector<int> topKFrequent(vector<int>& nums, int k) {
    unordered_map<int, int> freq;
    for (int n : nums) freq[n]++;

    vector<vector<int>> buckets(nums.size() + 1);
    for (auto& [val, cnt] : freq) buckets[cnt].push_back(val);

    vector<int> res;
    for (int i = buckets.size() - 1; i >= 0 && (int)res.size() < k; i--) {
        for (int val : buckets[i]) {
            res.push_back(val);
            if ((int)res.size() == k) break;
        }
    }
    return res;
}

int main() {
    vector<int> t1 = {1,1,1,2,2,3};
    auto r1 = topKFrequent(t1, 2);
    assert(r1.size() == 2 && r1[0] == 1);

    vector<int> t2 = {1};
    auto r2 = topKFrequent(t2, 1);
    assert(r2 == vector<int>{1});

    vector<int> t3 = {4,4,4,3,3,2,2,2,2};
    auto r3 = topKFrequent(t3, 1);
    assert(r3[0] == 2);

    cout << "All tests passed!" << endl;
    return 0;
}
