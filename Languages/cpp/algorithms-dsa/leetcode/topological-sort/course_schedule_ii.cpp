/**
 * LeetCode 210: Course Schedule II
 * Topic: Topological Sort
 * Difficulty: Medium
 *
 * Return the ordering of courses you should take (topological order).
 * Kahn's algorithm: BFS with in-degree tracking.
 * Time: O(V + E), Space: O(V + E)
 */
#include <iostream>
#include <vector>
#include <queue>
#include <cassert>
using namespace std;

vector<int> findOrder(int numCourses, vector<vector<int>>& prerequisites) {
    vector<vector<int>> graph(numCourses);
    vector<int> inDegree(numCourses, 0);

    for (auto& p : prerequisites) {
        graph[p[1]].push_back(p[0]);
        inDegree[p[0]]++;
    }

    queue<int> q;
    for (int i = 0; i < numCourses; i++) {
        if (inDegree[i] == 0) q.push(i);
    }

    vector<int> order;
    while (!q.empty()) {
        int course = q.front(); q.pop();
        order.push_back(course);
        for (int next : graph[course]) {
            if (--inDegree[next] == 0) {
                q.push(next);
            }
        }
    }

    return order.size() == (size_t)numCourses ? order : vector<int>{};
}

int main() {
    vector<vector<int>> p1 = {{1,0}};
    auto r1 = findOrder(2, p1);
    assert(r1.size() == 2 && r1[0] == 0 && r1[1] == 1);

    vector<vector<int>> p2 = {{1,0},{2,0},{3,1},{3,2}};
    auto r2 = findOrder(4, p2);
    assert(r2.size() == 4);

    // Cycle: impossible
    vector<vector<int>> p3 = {{1,0},{0,1}};
    auto r3 = findOrder(2, p3);
    assert(r3.empty());

    cout << "All tests passed!" << endl;
    return 0;
}
