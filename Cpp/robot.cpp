#include <iostream>
#include <vector>
#include <map>
#include <algorithm>

using namespace std;

int maxItemsCollected(vector<vector<int>> &items, int startPos, int k) {
    map<int, int> itemMap;

    // Store items in a map where key=position, value=sum of items
    for (auto &item : items) {
        itemMap[item[0]] += item[1];
    }

    // Convert map to sorted vector of positions
    vector<int> positions;
    for (auto &p : itemMap) {
        positions.push_back(p.first);
    }

    int n = positions.size();

    // Compute prefix sum for fast range queries
    vector<int> prefixSum(n + 1, 0);
    for (int i = 0; i < n; ++i) {
        prefixSum[i + 1] = prefixSum[i] + itemMap[positions[i]];
    }

    int maxCollected = 0;

    // Try all possible leftmost positions within k steps
    for (int left = 0; left < n; ++left) {
        if (abs(positions[left] - startPos) > k) continue;

        int remainingSteps = k - abs(positions[left] - startPos);

        // Find the farthest right position reachable within remainingSteps
        int right = upper_bound(positions.begin(), positions.end(), positions[left] + remainingSteps) - positions.begin() - 1;

        int collected = prefixSum[right + 1] - prefixSum[left];
        maxCollected = max(maxCollected, collected);
    }

    return maxCollected;
}

int main() {
    vector<vector<int>> items = {{2, 8}, {6, 3}, {8, 6}};
    int startPos = 6, k = 4;
    
    cout << "Max Items Collected: " << maxItemsCollected(items, startPos, k) << endl;
    
    return 0;
}

