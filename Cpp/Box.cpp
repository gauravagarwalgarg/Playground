#include <iostream>
#include <vector>
using namespace std;

class BrickWall {
private:
    static int leastNumber(vector<vector<int>>& array, int width) {
        vector<int> hash(width + 1, 0); // Hash array to store occurrences of edges
        int max = 0;

        // Iterating over input array
        for (size_t i = 0; i < array.size(); i++) {
            int sum = 0;
            for (size_t j = 0; j < array[i].size() - 1; j++) { // Exclude last brick
                sum = sum + array[i][j]; // Compute prefix sum (edge position)
                hash[sum] += 1;     // Count occurrences of each edge
            }
        }

        // Finding the max occurrences of edges
        for (int i = 0; i < width; i++) { // Exclude width itself
            if (hash[i] > max) {
                max = hash[i];
            }
        }

        return array.size() - max;
    }

public:
    static void run() {
        vector<vector<int>> array = {
            //{3, 3, 2, 1},
            //{5, 3},
            //{4, 4}
             {3, 5, 1, 1},
             {2, 3, 3, 2},
             {5, 5},
             {4, 4, 2},
             {1, 3, 3, 3},
             {1, 1, 6, 1, 1}};
        int width = 10;
        cout << leastNumber(array, width) << endl; // Expected output: 1
    }
};

int main() {
    BrickWall::run();
    return 0;
}

