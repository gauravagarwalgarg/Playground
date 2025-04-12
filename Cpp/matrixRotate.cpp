#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

void transpose(vector<vector<int>>& matrix) {
    int size = matrix.size();
    for (int i = 0; i < size; ++i) {
        for (int j = i + 1; j < size; ++j) {
            swap(matrix[i][j], matrix[j][i]);
        }
    }
}

void reverseRows(vector<vector<int>>& matrix) {
    for (auto& row : matrix) {
        reverse(row.begin(), row.end());
    }
}

void reverseColumns(vector<vector<int>>& matrix) {
    int size = matrix.size();
    std::cout << "size of reverseColumns: " << size << std::endl;
    for (int col = 0; col < size; ++col) {
        for (int row = 0; row < size / 2; ++row) {
            swap(matrix[row][col], matrix[size - row - 1][col]);
        }
    }
}

void rotateMatrix(vector<vector<int>>& matrix, int degrees, bool clockwise = true) {
    degrees = (clockwise ? degrees : 360 - degrees) % 360;  // Convert counterclockwise to equivalent clockwise

    if (degrees == 90) {
        reverseColumns(matrix);
        transpose(matrix);
    } 
    else if (degrees == 180) {
        reverseRows(matrix);
        reverseColumns(matrix);
    } 
    else if (degrees == 270) {
        transpose(matrix);
        reverseColumns(matrix);
    }
}

void printMatrix(const vector<vector<int>>& matrix) {
    for (const auto& row : matrix) {
        for (int num : row) {
            cout << num << " ";
        }
        cout << endl;
    }
}

int main() {
    vector<vector<int>> matrix = {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9}
    };

    int degrees;
    bool clockwise;
    cout << "Enter rotation degrees (90, 180, 270, 360): ";
    cin >> degrees;
    cout << "Clockwise? (1 for Yes, 0 for No): ";
    cin >> clockwise;

    rotateMatrix(matrix, degrees, clockwise);
    printMatrix(matrix);

    return 0;
}

