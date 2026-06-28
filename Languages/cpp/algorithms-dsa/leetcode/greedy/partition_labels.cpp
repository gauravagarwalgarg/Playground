/**
 * LeetCode 763: Partition Labels
 * Topic: Greedy
 * Difficulty: Medium
 *
 * Partition a string so each letter appears in at most one part.
 * Track last occurrence of each character, extend partition boundary.
 * Time: O(n), Space: O(1)
 */
#include <iostream>
#include <vector>
#include <string>
#include <cassert>
using namespace std;

vector<int> partitionLabels(string s) {
    vector<int> last(26, 0);
    for (int i = 0; i < (int)s.size(); i++)
        last[s[i] - 'a'] = i;

    vector<int> res;
    int start = 0, end = 0;
    for (int i = 0; i < (int)s.size(); i++) {
        end = max(end, last[s[i] - 'a']);
        if (i == end) {
            res.push_back(end - start + 1);
            start = i + 1;
        }
    }
    return res;
}

int main() {
    auto r1 = partitionLabels("ababcbacadefegdehijhklij");
    assert(r1 == vector<int>({9, 7, 8}));

    auto r2 = partitionLabels("eccbbbbdec");
    assert(r2 == vector<int>({10}));

    auto r3 = partitionLabels("abc");
    assert(r3 == vector<int>({1, 1, 1}));

    cout << "All tests passed!" << endl;
    return 0;
}
