#include <iostream>
#include <vector>
#include <sstream>
#include <fstream>
#include <map>
#include <algorithm>

using namespace std;

int minBrokenBars(const vector<vector<int>>& chocolateRows) {
    if (chocolateRows.empty()) return 0;

    map<int, int> cutPositions;
    int totalRows = chocolateRows.size();
    
    for (const auto& row : chocolateRows) {
        int position = 0;
        for (size_t i = 0; i < row.size() - 1; i++) { // Ignore last position
            position += row[i];
            cutPositions[position]++; // Count the number of times we can cut cleanly
        }
    }

    int minBroken = totalRows; // Assume worst case (all bars break)
    for (const auto& cut : cutPositions) {
        minBroken = min(minBroken, totalRows - cut.second); // Minimize broken bars
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

