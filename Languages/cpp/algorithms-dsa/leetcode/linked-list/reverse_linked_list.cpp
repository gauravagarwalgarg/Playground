/*
 * LeetCode 206 - Reverse Linked List
 * Topic: Linked List
 * Difficulty: Easy
 *
 * Iterative pointer swap: prev, curr, next.
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

ListNode* reverseList(ListNode* head) {
    ListNode* prev = nullptr;
    while (head) {
        ListNode* next = head->next;
        head->next = prev;
        prev = head;
        head = next;
    }
    return prev;
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
    assert(equals(reverseList(buildList({1,2,3,4,5})), {5,4,3,2,1}));
    assert(equals(reverseList(buildList({1,2})), {2,1}));
    assert(reverseList(nullptr) == nullptr);

    cout << "All tests passed!" << endl;
    return 0;
}
