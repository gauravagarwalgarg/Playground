/**
 * LeetCode 136: Single Number
 * Topic: Bit Manipulation
 * Difficulty: Easy
 *
 * Every element appears twice except one. Find the single element.
 * XOR all elements: duplicates cancel, leaving the single number.
 * Time: O(n), Space: O(1)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

int singleNumber(vector<int>& nums) {
    int result = 0;
    for (int n : nums) result ^= n;
    return result;
}

int main() {
    vector<int> v1 = {2, 2, 1};
    assert(singleNumber(v1) == 1);

    vector<int> v2 = {4, 1, 2, 1, 2};
    assert(singleNumber(v2) == 4);

    vector<int> v3 = {1};
    assert(singleNumber(v3) == 1);

    cout << "All tests passed!" << endl;
    return 0;
}
