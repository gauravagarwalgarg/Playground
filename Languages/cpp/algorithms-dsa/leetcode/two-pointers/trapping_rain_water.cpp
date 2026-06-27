/*
 * LeetCode 42 - Trapping Rain Water
 * Topic: Two Pointers
 * Difficulty: Hard
 *
 * Two pointers from both ends; water at each position = min(leftMax, rightMax) - height.
 * Time: O(n)
 * Space: O(1)
 */
#include <iostream>
#include <vector>
#include <cassert>
using namespace std;

int trap(vector<int>& height) {
    int left = 0, right = height.size() - 1;
    int leftMax = 0, rightMax = 0, water = 0;
    while (left < right) {
        if (height[left] < height[right]) {
            leftMax = max(leftMax, height[left]);
            water += leftMax - height[left];
            left++;
        } else {
            rightMax = max(rightMax, height[right]);
            water += rightMax - height[right];
            right--;
        }
    }
    return water;
}

int main() {
    vector<int> t1 = {0,1,0,2,1,0,1,3,2,1,2,1};
    assert(trap(t1) == 6);

    vector<int> t2 = {4,2,0,3,2,5};
    assert(trap(t2) == 9);

    vector<int> t3 = {1,2,3,4};
    assert(trap(t3) == 0);

    vector<int> t4 = {3,0,3};
    assert(trap(t4) == 3);

    cout << "All tests passed!" << endl;
    return 0;
}
