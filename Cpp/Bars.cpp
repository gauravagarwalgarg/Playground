#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

int minBrokenBars(vector<vector<int>>& plate) {
    unordered_map<int, int> cutPositions;
    int n = plate.size();
    
    for (const auto& row : plate) {
        int sum = 0;
        for (int piece : row) {
            sum += piece;
            cutPositions[sum]++;  // Mark where a cut could be made
        }
    }
    
    int totalRows = plate.size();
    int totalCols = 0;
    for (int piece : plate[0]) totalCols += piece;
    
    int minBroken = totalRows;  // Worst case: breaking all rows
    for (const auto& [cutPos, count] : cutPositions) {
        if (cutPos < totalCols) {
            minBroken = min(minBroken, totalRows - count);
        }
    }
    
    return minBroken;
}

int main() {
    int n;
    cin >> n;
    
    vector<vector<int>> plate(n);
    
    for (int i = 0; i < n; i++) {
        int bars;
        cin >> bars;
        plate[i].resize(bars);
        for (int j = 0; j < bars; j++) {
            cin >> plate[i][j];
        }
    }
    
    cout << minBrokenBars(plate) << endl;
    
    return 0;
}

