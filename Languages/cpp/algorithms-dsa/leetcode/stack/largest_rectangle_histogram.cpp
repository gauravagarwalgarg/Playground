/*
 * LeetCode 84 - Largest Rectangle in Histogram
 * Topic: Stack
 * Difficulty: Hard
 *
 * Monotonic increasing stack. When a shorter bar is found, pop and compute area.
 * Time: O(n)
 * Space: O(n)
 */
#include <iostream>
#include <vector>
#include <stack>
#include <cassert>
using namespace std;

int largestRectangleArea(vector<int>& heights) {
    stack<int> st;
    int maxArea = 0, n = heights.size();
    for (int i = 0; i <= n; i++) {
        int h = (i == n) ? 0 : heights[i];
        while (!st.empty() && h < heights[st.top()]) {
            int height = heights[st.top()]; st.pop();
            int width = st.empty() ? i : i - st.top() - 1;
            maxArea = max(maxArea, height * width);
        }
        st.push(i);
    }
    return maxArea;
}

int main() {
    vector<int> t1 = {2,1,5,6,2,3};
    assert(largestRectangleArea(t1) == 10);

    vector<int> t2 = {2,4};
    assert(largestRectangleArea(t2) == 4);

    vector<int> t3 = {1,1,1,1};
    assert(largestRectangleArea(t3) == 4);

    vector<int> t4 = {6,2,5,4,5,1,6};
    assert(largestRectangleArea(t4) == 12);

    cout << "All tests passed!" << endl;
    return 0;
}
