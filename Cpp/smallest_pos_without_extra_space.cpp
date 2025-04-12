#include <iostream>
#include <vector>

using namespace std;

int smallest_pos_number(vector<int>& nums) {
    int n = nums.size();

    // Step 1: Place each number at its correct index
    for (int i = 0; i < n; i++) {
        while (nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i] - 1]) {
            swap(nums[i], nums[nums[i] - 1]);
        }
    }

    // Step 2: Find the missing number
    for (int i = 0; i < n; i++) {
        if (nums[i] != i + 1) {
            return i + 1;
        }
    }

    return n + 1;
}

int main() {
    vector<int> nums = {2, -3, 4, 1, 1, 7};
    cout << smallest_pos_number(nums) << endl;
    return 0;
}

