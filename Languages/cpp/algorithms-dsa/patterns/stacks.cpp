/**
 * Pattern: Stacks (Monotonic Stack & Classic Stack Problems)
 *
 * Key techniques covered:
 * 1. Next Greater Element (monotonic stack decreasing)
 * 2. Daily Temperatures (monotonic stack indices)
 * 3. Valid Parentheses (stack matching)
 * 4. Largest Rectangle in Histogram (monotonic stack increasing)
 *
 * Core insight: A monotonic stack maintains elements in sorted order.
 * - Decreasing stack: finds the next greater element to the right.
 * - Increasing stack: finds the next smaller element to the right.
 *
 * Compile: g++ -std=c++17 -o test stacks.cpp && ./test
 */
#include <iostream>
#include <vector>
#include <stack>
#include <string>
#include <cassert>
using namespace std;

// ─────────────────────────────────────────────────────────────────────────────
// 1. Next Greater Element
//    For each element, find the first element to its right that is greater.
//    Strategy: Traverse from right to left, maintain a decreasing stack.
//    Time: O(n), Space: O(n)
// ─────────────────────────────────────────────────────────────────────────────
vector<int> nextGreaterElement(vector<int>& nums) {
    int n = nums.size();
    vector<int> result(n, -1);
    stack<int> st; // Stores values (decreasing from bottom to top)

    // Traverse from right to left
    for (int i = n - 1; i >= 0; i--) {
        // Pop elements that are not greater than current
        while (!st.empty() && st.top() <= nums[i]) {
            st.pop();
        }
        // If stack is not empty, top is the next greater element
        if (!st.empty()) {
            result[i] = st.top();
        }
        st.push(nums[i]);
    }
    return result;
}

// ─────────────────────────────────────────────────────────────────────────────
// 2. Daily Temperatures
//    For each day, find how many days until a warmer temperature.
//    Strategy: Monotonic stack storing indices, traverse left to right.
//    Time: O(n), Space: O(n)
// ─────────────────────────────────────────────────────────────────────────────
vector<int> dailyTemperatures(vector<int>& temps) {
    int n = temps.size();
    vector<int> result(n, 0);
    stack<int> st; // Stores indices

    for (int i = 0; i < n; i++) {
        // Pop indices whose temperatures are less than current
        while (!st.empty() && temps[st.top()] < temps[i]) {
            int prev = st.top();
            st.pop();
            result[prev] = i - prev;
        }
        st.push(i);
    }
    return result;
}

// ─────────────────────────────────────────────────────────────────────────────
// 3. Valid Parentheses
//    Check if a string of brackets is valid.
//    Strategy: Push opening brackets, pop and match closing brackets.
//    Time: O(n), Space: O(n)
// ─────────────────────────────────────────────────────────────────────────────
bool isValid(string s) {
    stack<char> st;

    for (char c : s) {
        if (c == '(' || c == '{' || c == '[') {
            st.push(c);
        } else {
            if (st.empty()) return false;
            char top = st.top();
            if ((c == ')' && top != '(') ||
                (c == '}' && top != '{') ||
                (c == ']' && top != '[')) {
                return false;
            }
            st.pop();
        }
    }
    return st.empty();
}

// ─────────────────────────────────────────────────────────────────────────────
// 4. Largest Rectangle in Histogram
//    Find the area of the largest rectangle that fits in the histogram.
//    Strategy: Maintain an increasing stack of indices. When we encounter
//    a bar shorter than stack top, we calculate areas using popped bars.
//    Time: O(n), Space: O(n)
// ─────────────────────────────────────────────────────────────────────────────
int largestRectangleArea(vector<int>& heights) {
    int n = heights.size();
    stack<int> st; // Indices of bars in increasing height order
    int maxArea = 0;

    for (int i = 0; i <= n; i++) {
        // Use 0 height as sentinel to flush remaining bars
        int currHeight = (i == n) ? 0 : heights[i];

        while (!st.empty() && heights[st.top()] > currHeight) {
            int h = heights[st.top()];
            st.pop();
            // Width: if stack empty, width = i; else width = i - st.top() - 1
            int w = st.empty() ? i : i - st.top() - 1;
            maxArea = max(maxArea, h * w);
        }
        st.push(i);
    }
    return maxArea;
}

// ─────────────────────────────────────────────────────────────────────────────
int main() {
    // Test 1: Next Greater Element
    {
        vector<int> nums = {4, 5, 2, 25};
        auto result = nextGreaterElement(nums);
        assert(result == vector<int>({5, 25, 25, -1}));

        vector<int> nums2 = {13, 7, 6, 12};
        auto result2 = nextGreaterElement(nums2);
        assert(result2 == vector<int>({-1, 12, 12, -1}));
    }

    // Test 2: Daily Temperatures
    {
        vector<int> temps = {73, 74, 75, 71, 69, 72, 76, 73};
        auto result = dailyTemperatures(temps);
        assert(result == vector<int>({1, 1, 4, 2, 1, 1, 0, 0}));
    }

    // Test 3: Valid Parentheses
    {
        assert(isValid("()") == true);
        assert(isValid("()[]{}") == true);
        assert(isValid("(]") == false);
        assert(isValid("([)]") == false);
        assert(isValid("{[]}") == true);
        assert(isValid("") == true);
    }

    // Test 4: Largest Rectangle in Histogram
    {
        vector<int> h1 = {2, 1, 5, 6, 2, 3};
        assert(largestRectangleArea(h1) == 10);

        vector<int> h2 = {2, 4};
        assert(largestRectangleArea(h2) == 4);

        vector<int> h3 = {1};
        assert(largestRectangleArea(h3) == 1);
    }

    cout << "All stack pattern tests passed!" << endl;
    return 0;
}
