/*
 * LeetCode 169 - Majority Element
 * Topic: Arrays & Hashing
 * Difficulty: Easy
 *
 * Boyer-Moore Voting Algorithm: maintain a candidate and count.
 * Time: O(n)
 * Space: O(1)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

int majorityElement(vector<int>& nums) {
    int candidate = 0, count = 0;
    for (int n : nums) {
        if (count == 0) candidate = n;
        count += (n == candidate) ? 1 : -1;
    }
    return candidate;
}

int main() {
    vector<int> t1 = {3,2,3};
    assert(majorityElement(t1) == 3);

    vector<int> t2 = {2,2,1,1,1,2,2};
    assert(majorityElement(t2) == 2);

    vector<int> t3 = {1};
    assert(majorityElement(t3) == 1);

    vector<int> t4 = {6,6,6,7,7};
    assert(majorityElement(t4) == 6);

    cout << "All tests passed!" << endl;
    return 0;
}
