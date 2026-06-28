/*
 * LeetCode 21 - Merge Two Sorted Lists
 * Topic: Linked List
 * Difficulty: Easy
 *
 * Iterative merge using a dummy head node.
 * Time: O(n + m)
 * Space: O(1)
 */
#include <iostream>
#include <cassert>
using namespace std;

struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(nullptr) {}
};

ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
    ListNode dummy(0);
    ListNode* curr = &dummy;
    while (l1 && l2) {
        if (l1->val <= l2->val) { curr->next = l1; l1 = l1->next; }
        else { curr->next = l2; l2 = l2->next; }
        curr = curr->next;
    }
    curr->next = l1 ? l1 : l2;
    return dummy.next;
}

ListNode* buildList(initializer_list<int> vals) {
    ListNode dummy(0);
    ListNode* curr = &dummy;
    for (int v : vals) { curr->next = new ListNode(v); curr = curr->next; }
    return dummy.next;
}

bool equals(ListNode* a, initializer_list<int> vals) {
    for (int v : vals) { if (!a || a->val != v) return false; a = a->next; }
    return a == nullptr;
}

int main() {
    assert(equals(mergeTwoLists(buildList({1,2,4}), buildList({1,3,4})), {1,1,2,3,4,4}));
    assert(equals(mergeTwoLists(nullptr, nullptr), {}));
    assert(equals(mergeTwoLists(nullptr, buildList({0})), {0}));

    cout << "All tests passed!" << endl;
    return 0;
}
