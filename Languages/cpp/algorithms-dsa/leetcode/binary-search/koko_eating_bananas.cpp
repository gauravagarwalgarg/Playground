/*
 * LeetCode 875 - Koko Eating Bananas
 * Topic: Binary Search
 * Difficulty: Medium
 *
 * Binary search on the answer space (eating speed). Check if Koko can finish in h hours.
 * Time: O(n * log m) where m = max pile size
 * Space: O(1)
 */
#include <iostream>
#include <vector>
#include <algorithm>
#include <cmath>
#include <cassert>
using namespace std;

int minEatingSpeed(vector<int>& piles, int h) {
    int lo = 1, hi = *max_element(piles.begin(), piles.end());
    while (lo < hi) {
        int mid = lo + (hi - lo) / 2;
        long hours = 0;
        for (int p : piles) hours += (p + mid - 1) / mid;
        if (hours <= h) hi = mid;
        else lo = mid + 1;
    }
    return lo;
}

int main() {
    vector<int> t1 = {3,6,7,11};
    assert(minEatingSpeed(t1, 8) == 4);

    vector<int> t2 = {30,11,23,4,20};
    assert(minEatingSpeed(t2, 5) == 30);

    vector<int> t3 = {30,11,23,4,20};
    assert(minEatingSpeed(t3, 6) == 23);

    cout << "All tests passed!" << endl;
    return 0;
}
