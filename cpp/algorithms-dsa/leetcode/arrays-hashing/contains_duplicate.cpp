/**
 * LeetCode 217: Contains Duplicate
 * Topic: Arrays & Hashing
 * Difficulty: Easy
 *
 * Return true if any value appears at least twice.
 * Time: O(n), Space: O(n)
 */
#include <iostream>
#include <vector>
#include <unordered_set>
#include <cassert>
using namespace std;

bool containsDuplicate(vector<int>& nums) {
    unordered_set<int> seen;
    for (int n : nums) {
        if (seen.count(n)) return true;
        seen.insert(n);
    }
    return false;
}

int main() {
    vector<int> v1 = {1, 2, 3, 1};
    assert(containsDuplicate(v1) == true);

    vector<int> v2 = {1, 2, 3, 4};
    assert(containsDuplicate(v2) == false);

    vector<int> v3 = {1, 1, 1, 3, 3, 4, 3, 2, 4, 2};
    assert(containsDuplicate(v3) == true);

    cout << "All tests passed!" << endl;
    return 0;
}
