/**
 * LeetCode 11: Container With Most Water
 * Topic: Two Pointers
 * Difficulty: Medium
 *
 * You are given an integer array height of length n. There are n vertical
 * lines drawn such that the two endpoints of the ith line are (i, 0) and
 * (i, height[i]).
 *
 * Find two lines that together with the x-axis form a container, such that
 * the container contains the most water.
 *
 * Approach: Two pointers starting from both ends. Calculate area, then move
 * the pointer pointing to the shorter line inward.
 *
 * Time: O(n), Space: O(1)
 */
#include <iostream>
#include <vector>
#include <algorithm>
#include <cassert>
using namespace std;

int maxArea(vector<int>& height) {
    int left = 0, right = height.size() - 1;
    int maxWater = 0;

    while (left < right) {
        int width = right - left;
        int h = min(height[left], height[right]);
        maxWater = max(maxWater, width * h);

        if (height[left] < height[right])
            left++;
        else
            right--;
    }

    return maxWater;
}

int main() {
    vector<int> v1 = {1, 8, 6, 2, 5, 4, 8, 3, 7};
    assert(maxArea(v1) == 49);

    vector<int> v2 = {1, 1};
    assert(maxArea(v2) == 1);

    vector<int> v3 = {4, 3, 2, 1, 4};
    assert(maxArea(v3) == 16);

    vector<int> v4 = {1, 2, 1};
    assert(maxArea(v4) == 2);

    vector<int> v5 = {2, 3, 4, 5, 18, 17, 6};
    assert(maxArea(v5) == 17);

    cout << "All tests passed!" << endl;
    return 0;
}
