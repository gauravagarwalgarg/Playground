/**
 * LeetCode 721: Accounts Merge
 * Topic: Union-Find
 * Difficulty: Medium
 *
 * Given accounts where each has a name and emails, merge accounts that
 * share at least one email. Union-Find on email ownership.
 * Time: O(n*k * α(n*k)) where k = avg emails per account, Space: O(n*k)
 */
#include <iostream>
#include <vector>
#include <string>
#include <unordered_map>
#include <set>
#include <algorithm>
#include <cassert>
using namespace std;

class UnionFind {
    vector<int> parent;
public:
    UnionFind(int n) : parent(n) {
        for (int i = 0; i < n; i++) parent[i] = i;
    }
    int find(int x) {
        if (parent[x] != x) parent[x] = find(parent[x]);
        return parent[x];
    }
    void unite(int x, int y) {
        parent[find(x)] = find(y);
    }
};

vector<vector<string>> accountsMerge(vector<vector<string>>& accounts) {
    int n = accounts.size();
    UnionFind uf(n);
    unordered_map<string, int> emailToId;

    for (int i = 0; i < n; i++) {
        for (int j = 1; j < (int)accounts[i].size(); j++) {
            string& email = accounts[i][j];
            if (emailToId.count(email)) {
                uf.unite(i, emailToId[email]);
            } else {
                emailToId[email] = i;
            }
        }
    }

    unordered_map<int, set<string>> merged;
    for (int i = 0; i < n; i++) {
        int root = uf.find(i);
        for (int j = 1; j < (int)accounts[i].size(); j++)
            merged[root].insert(accounts[i][j]);
    }

    vector<vector<string>> res;
    for (auto& [root, emails] : merged) {
        vector<string> account = {accounts[root][0]};
        for (auto& e : emails) account.push_back(e);
        res.push_back(account);
    }
    return res;
}

int main() {
    vector<vector<string>> accounts = {
        {"John", "johnsmith@mail.com", "john_newyork@mail.com"},
        {"John", "johnsmith@mail.com", "john00@mail.com"},
        {"Mary", "mary@mail.com"},
        {"John", "johnnybravo@mail.com"}
    };
    auto res = accountsMerge(accounts);
    assert(res.size() == 3);

    // Find the merged John account (should have 3 emails)
    bool foundMerged = false;
    for (auto& acc : res) {
        if (acc.size() == 4 && acc[0] == "John") foundMerged = true;
    }
    assert(foundMerged);

    cout << "All tests passed!" << endl;
    return 0;
}
