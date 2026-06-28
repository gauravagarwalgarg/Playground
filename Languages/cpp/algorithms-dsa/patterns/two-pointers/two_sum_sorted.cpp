/**
 * Pattern: Two Pointers
 * Problem: Two Sum II - Input Array Is Sorted (LeetCode 167)
 *
 * Given a 1-indexed sorted array, find two numbers that add up to target.
 * Time: O(n), Space: O(1)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

vector<int> twoSum(vector<int>& numbers, int target) {
    int left = 0, right = numbers.size() - 1;
    while (left < right) {
        int sum = numbers[left] + numbers[right];
        if (sum == target) return {left + 1, right + 1};
        else if (sum < target) left++;
        else right--;
    }
    return {};
}

int main() {
    vector<int> v1 = {2, 7, 11, 15};
    assert(twoSum(v1, 9) == vector<int>({1, 2}));

    vector<int> v2 = {2, 3, 4};
    assert(twoSum(v2, 6) == vector<int>({1, 3}));

    vector<int> v3 = {-1, 0};
    assert(twoSum(v3, -1) == vector<int>({1, 2}));

    cout << "All tests passed!" << endl;
    return 0;
}
