#include <iostream>
#include <vector>
#include <cmath>

using namespace std;

int difference_calculator(int N, vector<int>& arr) {
    int alpha = 0, beta = 0;

    for (int i = 0; i < N; i++) {
        int count = 0, j = i;

        while (j < N && arr[j] == arr[i]) {
            count++;
            j++;
        }

        if (count == arr[i]) {
            alpha++;
        }

        if ((i + 1) == arr[i] && count == arr[i]) {
            beta++;
        }

        i = j - 1;
    }

    return abs(alpha - beta);
}

int main() {
    vector<int> arr = {2, 2, 2, 4, 4, 4, 4, 1, 2, 2}; 
    int N = arr.size();
    cout << "Difference (|alpha - beta|): " << difference_calculator(N, arr) << endl;
    return 0;
}

