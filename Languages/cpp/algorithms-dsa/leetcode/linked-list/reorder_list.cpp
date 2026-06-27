/*
 * LeetCode 143 - Reorder List
 * Topic: Linked List
 * Difficulty: Medium
 *
 * Find middle, reverse second half, merge alternating.
 * Time: O(n)
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

void reorderList(ListNode* head) {
    if (!head || !head->next) return;
    // Find middle
    ListNode *slow = head, *fast = head;
    while (fast->next && fast->next->next) { slow = slow->next; fast = fast->next->next; }
    // Reverse second half
    ListNode* prev = nullptr;
    ListNode* curr = slow->next;
    slow->next = nullptr;
    while (curr) { ListNode* nxt = curr->next; curr->next = prev; prev = curr; curr = nxt; }
    // Merge two halves
    ListNode *first = head, *second = prev;
    while (second) {
        ListNode *tmp1 = first->next, *tmp2 = second->next;
        first->next = second; second->next = tmp1;
        first = tmp1; second = tmp2;
    }
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
    ListNode* l1 = buildList({1,2,3,4});
    reorderList(l1);
    assert(equals(l1, {1,4,2,3}));

    ListNode* l2 = buildList({1,2,3,4,5});
    reorderList(l2);
    assert(equals(l2, {1,5,2,4,3}));

    ListNode* l3 = buildList({1});
    reorderList(l3);
    assert(equals(l3, {1}));

    cout << "All tests passed!" << endl;
    return 0;
}
