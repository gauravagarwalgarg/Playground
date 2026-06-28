/**
 * LeetCode 70: Climbing Stairs
 * Topic: Dynamic Programming
 * Difficulty: Easy
 *
 * You can climb 1 or 2 steps. How many distinct ways to reach the top?
 * Time: O(n), Space: O(1)
 */
#include <iostream>
#include <cassert>
using namespace std;

int climbStairs(int n) {
    if (n <= 2) return n;
    int prev2 = 1, prev1 = 2;
    for (int i = 3; i <= n; i++) {
        int curr = prev1 + prev2;
        prev2 = prev1;
        prev1 = curr;
    }
    return prev1;
}

int main() {
    assert(climbStairs(1) == 1);
    assert(climbStairs(2) == 2);
    assert(climbStairs(3) == 3);
    assert(climbStairs(5) == 8);
    assert(climbStairs(10) == 89);

    cout << "All tests passed!" << endl;
    return 0;
}
