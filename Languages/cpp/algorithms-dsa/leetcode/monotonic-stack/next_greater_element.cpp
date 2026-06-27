/**
 * LeetCode 496: Next Greater Element I
 * Topic: Monotonic Stack
 * Difficulty: Easy
 *
 * Find the next greater element for each element in nums1 from nums2.
 * Use a decreasing monotonic stack to process nums2 right-to-left (or left-to-right with pop).
 * Time: O(n + m), Space: O(n)
 */
#include <iostream>
#include <vector>
#include <stack>
#include <unordered_map>
#include <cassert>
using namespace std;

vector<int> nextGreaterElement(vector<int>& nums1, vector<int>& nums2) {
    unordered_map<int, int> nextGreater;
    stack<int> st; // decreasing stack

    for (int num : nums2) {
        while (!st.empty() && st.top() < num) {
            nextGreater[st.top()] = num;
            st.pop();
        }
        st.push(num);
    }

    vector<int> result;
    for (int num : nums1) {
        result.push_back(nextGreater.count(num) ? nextGreater[num] : -1);
    }
    return result;
}

int main() {
    vector<int> n1 = {4, 1, 2}, n2 = {1, 3, 4, 2};
    assert((nextGreaterElement(n1, n2) == vector<int>{-1, 3, -1}));

    vector<int> n3 = {2, 4}, n4 = {1, 2, 3, 4};
    assert((nextGreaterElement(n3, n4) == vector<int>{3, -1}));

    cout << "All tests passed!" << endl;
    return 0;
}
