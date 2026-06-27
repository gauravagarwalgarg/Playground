/**
 * LeetCode 621: Task Scheduler
 * Topic: Greedy
 * Difficulty: Medium
 *
 * Given tasks and cooldown n, find minimum intervals to finish all tasks.
 * Count frequencies; the most frequent task dictates idle slots.
 * Time: O(t) where t = number of tasks, Space: O(1) (26 letters max)
 */
#include <iostream>
#include <vector>
#include <cassert>
#include <algorithm>
using namespace std;

int leastInterval(vector<char>& tasks, int n) {
    vector<int> freq(26, 0);
    for (char c : tasks) freq[c - 'A']++;
    int maxFreq = *max_element(freq.begin(), freq.end());
    int maxCount = count(freq.begin(), freq.end(), maxFreq);
    int partitions = maxFreq - 1;
    int emptySlots = partitions * (n - (maxCount - 1));
    int available = (int)tasks.size() - maxFreq * maxCount;
    int idles = max(0, emptySlots - available);
    return (int)tasks.size() + idles;
}

int main() {
    vector<char> t1 = {'A','A','A','B','B','B'};
    assert(leastInterval(t1, 2) == 8);

    vector<char> t2 = {'A','A','A','B','B','B'};
    assert(leastInterval(t2, 0) == 6);

    vector<char> t3 = {'A','A','A','A','A','A','B','C','D','E','F','G'};
    assert(leastInterval(t3, 2) == 16);

    cout << "All tests passed!" << endl;
    return 0;
}
