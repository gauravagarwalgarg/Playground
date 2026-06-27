/**
 * LeetCode 269: Alien Dictionary
 * Topic: Topological Sort
 * Difficulty: Hard
 *
 * Given a sorted list of words in an alien language, derive the order of characters.
 * Build graph from adjacent word comparisons, then topological sort.
 * Time: O(C) where C = total chars across all words, Space: O(1) [26 letters max]
 */
#include <iostream>
#include <vector>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <queue>
#include <cassert>
using namespace std;

string alienOrder(vector<string>& words) {
    unordered_map<char, unordered_set<char>> graph;
    unordered_map<char, int> inDegree;

    // Initialize all characters
    for (auto& word : words)
        for (char c : word)
            inDegree[c]; // default 0

    // Build graph from adjacent word pairs
    for (int i = 0; i < (int)words.size() - 1; i++) {
        string& w1 = words[i];
        string& w2 = words[i + 1];
        int minLen = min(w1.size(), w2.size());

        // Check invalid case: "abc" before "ab"
        if (w1.size() > w2.size() && w1.substr(0, minLen) == w2.substr(0, minLen))
            return "";

        for (int j = 0; j < minLen; j++) {
            if (w1[j] != w2[j]) {
                if (!graph[w1[j]].count(w2[j])) {
                    graph[w1[j]].insert(w2[j]);
                    inDegree[w2[j]]++;
                }
                break;
            }
        }
    }

    // Kahn's BFS topological sort
    queue<char> q;
    for (auto& [c, deg] : inDegree)
        if (deg == 0) q.push(c);

    string result;
    while (!q.empty()) {
        char c = q.front(); q.pop();
        result += c;
        for (char next : graph[c]) {
            if (--inDegree[next] == 0)
                q.push(next);
        }
    }

    return result.size() == inDegree.size() ? result : "";
}

int main() {
    vector<string> w1 = {"wrt", "wrf", "er", "ett", "rftt"};
    string r1 = alienOrder(w1);
    assert(r1.size() == 5); // valid ordering exists

    vector<string> w2 = {"z", "x"};
    string r2 = alienOrder(w2);
    assert(r2 == "zx");

    vector<string> w3 = {"z", "x", "z"};
    string r3 = alienOrder(w3);
    assert(r3 == ""); // cycle

    cout << "All tests passed!" << endl;
    return 0;
}
