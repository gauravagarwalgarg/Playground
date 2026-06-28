/*
 * LeetCode 141 - Linked List Cycle
 * Topic: Linked List
 * Difficulty: Easy
 *
 * Floyd's cycle detection: fast pointer moves 2 steps, slow moves 1.
 * If they meet, there's a cycle.
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

bool hasCycle(ListNode* head) {
    ListNode* slow = head;
    ListNode* fast = head;
    while (fast && fast->next) {
        slow = slow->next;
        fast = fast->next->next;
        if (slow == fast) return true;
    }
    return false;
}

int main() {
    // Cycle: 3 -> 2 -> 0 -> -4 -> 2 (back to node at index 1)
    ListNode* n1 = new ListNode(3);
    ListNode* n2 = new ListNode(2);
    ListNode* n3 = new ListNode(0);
    ListNode* n4 = new ListNode(-4);
    n1->next = n2; n2->next = n3; n3->next = n4; n4->next = n2;
    assert(hasCycle(n1) == true);

    // No cycle: 1 -> 2
    ListNode* m1 = new ListNode(1);
    ListNode* m2 = new ListNode(2);
    m1->next = m2;
    assert(hasCycle(m1) == false);

    // Single node, no cycle
    assert(hasCycle(new ListNode(1)) == false);

    cout << "All tests passed!" << endl;
    return 0;
}
