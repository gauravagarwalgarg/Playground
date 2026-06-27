/*
 * LeetCode 133 - Clone Graph
 * Topic: Graphs
 * Difficulty: Medium
 *
 * DFS traversal with a hashmap to track visited/cloned nodes.
 * Time: O(V + E)
 * Space: O(V)
 */
#include <iostream>
#include <vector>
#include <unordered_map>
#include <cassert>
using namespace std;

struct Node {
    int val;
    vector<Node*> neighbors;
    Node(int v) : val(v) {}
};

unordered_map<Node*, Node*> visited;

Node* cloneGraph(Node* node) {
    if (!node) return nullptr;
    if (visited.count(node)) return visited[node];
    Node* clone = new Node(node->val);
    visited[node] = clone;
    for (auto* nb : node->neighbors)
        clone->neighbors.push_back(cloneGraph(nb));
    return clone;
}

int main() {
    // Build: 1 -- 2
    //        |    |
    //        4 -- 3
    Node* n1 = new Node(1); Node* n2 = new Node(2);
    Node* n3 = new Node(3); Node* n4 = new Node(4);
    n1->neighbors = {n2, n4}; n2->neighbors = {n1, n3};
    n3->neighbors = {n2, n4}; n4->neighbors = {n1, n3};

    visited.clear();
    Node* c1 = cloneGraph(n1);
    assert(c1 != n1);
    assert(c1->val == 1);
    assert(c1->neighbors.size() == 2);
    assert(c1->neighbors[0]->val == 2);

    // Empty graph
    visited.clear();
    assert(cloneGraph(nullptr) == nullptr);

    cout << "All tests passed!" << endl;
    return 0;
}
