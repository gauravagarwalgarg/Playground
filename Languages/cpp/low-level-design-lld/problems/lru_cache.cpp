/**
 * LLD Problem: LRU Cache (LeetCode #146)
 *
 * Design a data structure that follows the Least Recently Used (LRU) eviction
 * policy. Implement get(key) and put(key, value) in O(1) time.
 *
 * Approach: Doubly-linked list (for O(1) move/remove) + hash map (for O(1) lookup).
 * - get: move node to front, return value
 * - put: insert at front, evict from back if over capacity
 *
 * Time: O(1) for both get and put
 * Space: O(capacity)
 */
#include <iostream>
#include <unordered_map>
#include <cassert>
using namespace std;

struct Node {
    int key, value;
    Node* prev;
    Node* next;
    Node(int k, int v) : key(k), value(v), prev(nullptr), next(nullptr) {}
};

class LRUCache {
public:
    LRUCache(int capacity) : capacity_(capacity) {
        // Sentinel nodes to avoid null checks
        head_ = new Node(0, 0);
        tail_ = new Node(0, 0);
        head_->next = tail_;
        tail_->prev = head_;
    }

    ~LRUCache() {
        Node* curr = head_;
        while (curr) {
            Node* next = curr->next;
            delete curr;
            curr = next;
        }
    }

    int get(int key) {
        auto it = map_.find(key);
        if (it == map_.end()) return -1;
        Node* node = it->second;
        moveToFront(node);
        return node->value;
    }

    void put(int key, int value) {
        auto it = map_.find(key);
        if (it != map_.end()) {
            it->second->value = value;
            moveToFront(it->second);
        } else {
            if ((int)map_.size() == capacity_) {
                // Evict LRU (tail's prev)
                Node* lru = tail_->prev;
                removeNode(lru);
                map_.erase(lru->key);
                delete lru;
            }
            Node* node = new Node(key, value);
            addToFront(node);
            map_[key] = node;
        }
    }

private:
    int capacity_;
    Node* head_;
    Node* tail_;
    unordered_map<int, Node*> map_;

    void removeNode(Node* node) {
        node->prev->next = node->next;
        node->next->prev = node->prev;
    }

    void addToFront(Node* node) {
        node->next = head_->next;
        node->prev = head_;
        head_->next->prev = node;
        head_->next = node;
    }

    void moveToFront(Node* node) {
        removeNode(node);
        addToFront(node);
    }
};

int main() {
    LRUCache cache(2);

    cache.put(1, 1);
    cache.put(2, 2);
    assert(cache.get(1) == 1);

    cache.put(3, 3);  // evicts key 2
    assert(cache.get(2) == -1);

    cache.put(4, 4);  // evicts key 1
    assert(cache.get(1) == -1);
    assert(cache.get(3) == 3);
    assert(cache.get(4) == 4);

    // Update existing key
    cache.put(3, 30);
    assert(cache.get(3) == 30);

    cout << "All tests passed!" << endl;
    return 0;
}
