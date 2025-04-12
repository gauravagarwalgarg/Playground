#include <iostream>
#include <vector>
#include <sstream>
#include <fstream>
#include <algorithm>

using namespace std;

// Function to determine the minimum number of broken chocolate bars
int minBrokenBars(const vector<vector<int>>& chocolateRows) {
    if (chocolateRows.empty()) return 0;

    // Calculate prefix sums to determine possible cut positions
    vector<int> cutCounts;
    int maxWidth = 0;

    for (const auto& row : chocolateRows) {
        int sum = 0;
        vector<int> cutPositions;
        for (int pieces : row) {
            sum += pieces;
            cutPositions.push_back(sum);
        }
        maxWidth = max(maxWidth, sum);
        cutCounts.insert(cutCounts.end(), cutPositions.begin(), cutPositions.end());
    }

    // Count occurrences of each cut position
    sort(cutCounts.begin(), cutCounts.end());
    int minBroken = chocolateRows.size();
    for (size_t i = 0; i < cutCounts.size();) {
        int pos = cutCounts[i];
        int count = 0;
        while (i < cutCounts.size() && cutCounts[i] == pos) {
            count++;
            i++;
        }
        // Minimize the number of broken bars
        minBroken = min(minBroken, static_cast<int>(chocolateRows.size()) - count);
    }

    return minBroken;
}

int main() {
    // Read input from a file instead of using cin
    ifstream inputFile("input.txt");
    if (!inputFile) {
        cerr << "Error: Unable to open input file!" << endl;
        return 1;
    }

    string line;
    vector<vector<int>> chocolateRows;
    int n;

    // Read number of rows
    getline(inputFile, line);
    stringstream ss(line);
    ss >> n;

    // Read rows data
    for (int i = 0; i < n; i++) {
        int numBars;
        getline(inputFile, line);
        stringstream ssRow(line);
        ssRow >> numBars;

        vector<int> row;
        getline(inputFile, line);
        stringstream ssPieces(line);
        int piece;
        while (ssPieces >> piece) {
            row.push_back(piece);
        }
        chocolateRows.push_back(row);
    }

    inputFile.close();

    // Compute and print the minimum number of broken bars
    cout << minBrokenBars(chocolateRows) << endl;

    return 0;
}

