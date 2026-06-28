/*
 * LeetCode 560 - Subarray Sum Equals K
 * Topic: Arrays & Hashing
 * Difficulty: Medium
 *
 * Use prefix sum with hash map: count how many previous prefix sums equal (current - k).
 * Time: O(n)
 * Space: O(n)
 */
#include <iostream>
#include <vector>
#include <unordered_map>
#include <cassert>
using namespace std;

int subarraySum(vector<int>& nums, int k) {
    unordered_map<int, int> prefixCount;
    prefixCount[0] = 1;
    int sum = 0, count = 0;
    for (int n : nums) {
        sum += n;
        if (prefixCount.count(sum - k))
            count += prefixCount[sum - k];
        prefixCount[sum]++;
    }
    return count;
}

int main() {
    vector<int> t1 = {1,1,1};
    assert(subarraySum(t1, 2) == 2);

    vector<int> t2 = {1,2,3};
    assert(subarraySum(t2, 3) == 2);

    vector<int> t3 = {1,-1,0};
    assert(subarraySum(t3, 0) == 3);

    vector<int> t4 = {3,4,7,2,-3,1,4,2};
    assert(subarraySum(t4, 7) == 4);

    cout << "All tests passed!" << endl;
    return 0;
}
