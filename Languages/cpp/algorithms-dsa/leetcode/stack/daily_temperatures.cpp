/*
 * LeetCode 739 - Daily Temperatures
 * Topic: Stack
 * Difficulty: Medium
 *
 * Monotonic decreasing stack of indices. Pop when current temp is higher.
 * Time: O(n)
 * Space: O(n)
 */
#include <iostream>
#include <vector>
#include <stack>
#include <cassert>
using namespace std;

vector<int> dailyTemperatures(vector<int>& temps) {
    int n = temps.size();
    vector<int> res(n, 0);
    stack<int> st;
    for (int i = 0; i < n; i++) {
        while (!st.empty() && temps[i] > temps[st.top()]) {
            int idx = st.top(); st.pop();
            res[idx] = i - idx;
        }
        st.push(i);
    }
    return res;
}

int main() {
    vector<int> t1 = {73,74,75,71,69,72,76,73};
    assert(dailyTemperatures(t1) == (vector<int>{1,1,4,2,1,1,0,0}));

    vector<int> t2 = {30,40,50,60};
    assert(dailyTemperatures(t2) == (vector<int>{1,1,1,0}));

    vector<int> t3 = {30,60,90};
    assert(dailyTemperatures(t3) == (vector<int>{1,1,0}));

    cout << "All tests passed!" << endl;
    return 0;
}
