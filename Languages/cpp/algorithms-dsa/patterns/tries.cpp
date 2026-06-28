/**
 * Pattern: Tries (Prefix Trees)
 *
 * Key techniques covered:
 * 1. Trie Insert add a word character by character
 * 2. Trie Search check if a complete word exists
 * 3. Trie StartsWith check if any word starts with a given prefix
 *
 * Core insight: A Trie is a tree where each edge represents a character.
 * Each node has up to 26 children (for lowercase English letters).
 * A boolean flag marks the end of a complete word.
 *
 * Time complexity:
 * - Insert: O(L) where L = word length
 * - Search: O(L)
 * - StartsWith: O(L)
 * Space: O(total characters across all inserted words)
 *
 * Compile: g++ -std=c++17 -o test tries.cpp && ./test
 */
#include <iostream>
#include <string>
#include <cassert>
#include <memory>
using namespace std;

// ─────────────────────────────────────────────────────────────────────────────
// Trie Node
// Each node holds 26 children (one per lowercase letter) and an end-of-word flag.
// ─────────────────────────────────────────────────────────────────────────────
struct TrieNode {
    unique_ptr<TrieNode> children[26];
    bool isEndOfWord = false;
};

// ─────────────────────────────────────────────────────────────────────────────
// Trie Class
// ─────────────────────────────────────────────────────────────────────────────
class Trie {
    unique_ptr<TrieNode> root;

public:
    Trie() : root(make_unique<TrieNode>()) {}

    // Insert a word into the trie
    // Walk down the tree, creating nodes for characters that don't exist yet
    void insert(const string& word) {
        TrieNode* node = root.get();
        for (char c : word) {
            int idx = c - 'a';
            if (!node->children[idx]) {
                node->children[idx] = make_unique<TrieNode>();
            }
            node = node->children[idx].get();
        }
        node->isEndOfWord = true;
    }

    // Search for a complete word in the trie
    // Returns true only if the word exists AND is marked as end-of-word
    bool search(const string& word) {
        TrieNode* node = findNode(word);
        return node != nullptr && node->isEndOfWord;
    }

    // Check if any word in the trie starts with the given prefix
    // Returns true if we can traverse the entire prefix
    bool startsWith(const string& prefix) {
        return findNode(prefix) != nullptr;
    }

private:
    // Helper: traverse the trie following the given string
    // Returns nullptr if any character is missing
    TrieNode* findNode(const string& s) {
        TrieNode* node = root.get();
        for (char c : s) {
            int idx = c - 'a';
            if (!node->children[idx]) {
                return nullptr;
            }
            node = node->children[idx].get();
        }
        return node;
    }
};

// ─────────────────────────────────────────────────────────────────────────────
int main() {
    // Basic Trie operations
    {
        Trie trie;

        // Insert words
        trie.insert("apple");
        trie.insert("app");
        trie.insert("application");
        trie.insert("bat");
        trie.insert("ball");

        // Search exact word match
        assert(trie.search("apple") == true);
        assert(trie.search("app") == true);
        assert(trie.search("ap") == false);       // Prefix, not a complete word
        assert(trie.search("application") == true);
        assert(trie.search("bat") == true);
        assert(trie.search("bath") == false);     // Not inserted

        // StartsWith prefix match
        assert(trie.startsWith("app") == true);
        assert(trie.startsWith("appl") == true);
        assert(trie.startsWith("apple") == true);
        assert(trie.startsWith("b") == true);
        assert(trie.startsWith("bal") == true);
        assert(trie.startsWith("cat") == false);  // No word starts with "cat"
        assert(trie.startsWith("") == true);      // Empty prefix matches everything
    }

    // Edge cases
    {
        Trie trie;

        // Search in empty trie
        assert(trie.search("hello") == false);
        assert(trie.startsWith("h") == false);

        // Single character words
        trie.insert("a");
        assert(trie.search("a") == true);
        assert(trie.startsWith("a") == true);
        assert(trie.search("ab") == false);

        // Inserting same word twice is idempotent
        trie.insert("a");
        assert(trie.search("a") == true);
    }

    // Overlapping prefixes
    {
        Trie trie;
        trie.insert("the");
        trie.insert("there");
        trie.insert("their");
        trie.insert("them");

        assert(trie.search("the") == true);
        assert(trie.search("there") == true);
        assert(trie.search("their") == true);
        assert(trie.search("them") == true);
        assert(trie.search("then") == false);
        assert(trie.startsWith("the") == true);
        assert(trie.startsWith("ther") == true);
        assert(trie.startsWith("thei") == true);
    }

    cout << "All trie pattern tests passed!" << endl;
    return 0;
}
