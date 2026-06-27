/*
 * LeetCode 23 - Merge k Sorted Lists
 * Topic: Heap
 * Difficulty: Hard
 *
 * Use a min-heap to always pick the smallest head among k lists.
 * Time: O(n log k) where n = total nodes
 * Space: O(k)
 */
#include <iostream>
#include <vector>
#include <queue>
#include <cassert>
using namespace std;

struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(nullptr) {}
};

ListNode* mergeKLists(vector<ListNode*>& lists) {
    auto cmp = [](ListNode* a, ListNode* b) { return a->val > b->val; };
    priority_queue<ListNode*, vector<ListNode*>, decltype(cmp)> pq(cmp);
    for (auto* l : lists) if (l) pq.push(l);
    ListNode dummy(0);
    ListNode* curr = &dummy;
    while (!pq.empty()) {
        curr->next = pq.top(); pq.pop();
        curr = curr->next;
        if (curr->next) pq.push(curr->next);
    }
    return dummy.next;
}

ListNode* buildList(initializer_list<int> vals) {
    ListNode dummy(0); ListNode* c = &dummy;
    for (int v : vals) { c->next = new ListNode(v); c = c->next; }
    return dummy.next;
}

bool equals(ListNode* a, initializer_list<int> vals) {
    for (int v : vals) { if (!a || a->val != v) return false; a = a->next; }
    return a == nullptr;
}

int main() {
    vector<ListNode*> t1 = {buildList({1,4,5}), buildList({1,3,4}), buildList({2,6})};
    assert(equals(mergeKLists(t1), {1,1,2,3,4,4,5,6}));

    vector<ListNode*> t2 = {};
    assert(mergeKLists(t2) == nullptr);

    vector<ListNode*> t3 = {nullptr};
    assert(mergeKLists(t3) == nullptr);

    cout << "All tests passed!" << endl;
    return 0;
}
