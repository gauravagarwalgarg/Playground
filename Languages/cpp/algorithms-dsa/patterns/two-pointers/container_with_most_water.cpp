/**
 * Pattern: Two Pointers
 * Problem: Container With Most Water (LeetCode 11)
 *
 * Find two lines that together with x-axis form a container holding most water.
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
        int water = min(height[left], height[right]) * (right - left);
        maxWater = max(maxWater, water);
        if (height[left] < height[right]) left++;
        else right--;
    }
    return maxWater;
}

int main() {
    vector<int> h1 = {1, 8, 6, 2, 5, 4, 8, 3, 7};
    assert(maxArea(h1) == 49);

    vector<int> h2 = {1, 1};
    assert(maxArea(h2) == 1);

    cout << "All tests passed!" << endl;
    return 0;
}
