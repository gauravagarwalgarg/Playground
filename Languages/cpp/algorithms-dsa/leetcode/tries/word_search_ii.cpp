/**
 * LeetCode 212: Word Search II
 * Topic: Tries
 * Difficulty: Hard
 *
 * Given a board and a list of words, return all words that can be formed
 * by sequentially adjacent cells. Build a Trie from words, DFS on grid.
 * Time: O(m*n*4^L), Space: O(W*L) where W = words count, L = max word length
 */
#include <iostream>
#include <vector>
#include <string>
#include <algorithm>
#include <cassert>
using namespace std;

struct TrieNode {
    TrieNode* children[26];
    string word;
    TrieNode() : word("") {
        for (int i = 0; i < 26; i++) children[i] = nullptr;
    }
};

void insertWord(TrieNode* root, const string& w) {
    TrieNode* node = root;
    for (char c : w) {
        int idx = c - 'a';
        if (!node->children[idx]) node->children[idx] = new TrieNode();
        node = node->children[idx];
    }
    node->word = w;
}

void dfs(vector<vector<char>>& board, int i, int j, TrieNode* node, vector<string>& res) {
    if (i < 0 || i >= (int)board.size() || j < 0 || j >= (int)board[0].size()) return;
    char c = board[i][j];
    if (c == '#' || !node->children[c - 'a']) return;

    node = node->children[c - 'a'];
    if (!node->word.empty()) {
        res.push_back(node->word);
        node->word = ""; // avoid duplicates
    }

    board[i][j] = '#';
    dfs(board, i + 1, j, node, res);
    dfs(board, i - 1, j, node, res);
    dfs(board, i, j + 1, node, res);
    dfs(board, i, j - 1, node, res);
    board[i][j] = c;
}

vector<string> findWords(vector<vector<char>>& board, vector<string>& words) {
    TrieNode* root = new TrieNode();
    for (auto& w : words) insertWord(root, w);

    vector<string> res;
    for (int i = 0; i < (int)board.size(); i++)
        for (int j = 0; j < (int)board[0].size(); j++)
            dfs(board, i, j, root, res);

    sort(res.begin(), res.end());
    return res;
}

int main() {
    vector<vector<char>> board = {
        {'o','a','a','n'},
        {'e','t','a','e'},
        {'i','h','k','r'},
        {'i','f','l','v'}
    };
    vector<string> words = {"oath","pea","eat","rain"};
    auto r1 = findWords(board, words);
    assert(find(r1.begin(), r1.end(), "oath") != r1.end());
    assert(find(r1.begin(), r1.end(), "eat") != r1.end());
    assert(find(r1.begin(), r1.end(), "pea") == r1.end());

    cout << "All tests passed!" << endl;
    return 0;
}
