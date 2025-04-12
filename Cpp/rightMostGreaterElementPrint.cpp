#include <iostream>
#include <vector>
#include <stack>

using namespace std;

vector<int> rightMostGreaterElement(vector<int>& nums) {
    int n = nums.size();
    vector<int> res(n, -1); // Default all to -1
    stack<int> st; // Stack to store indices

    cout << "Processing array from right to left:\n";

    // Traverse from right to left
    for (int i = n - 1; i >= 0; i--) {
        cout << "\nChecking element nums[" << i << "] = " << nums[i] << endl;

        // Remove elements from stack that are smaller than or equal to nums[i]
        while (!st.empty() && nums[st.top()] <= nums[i]) {
            cout << "  Popping " << nums[st.top()] << " from stack as it's <= " << nums[i] << endl;
            st.pop();
        }

        // If stack is not empty, the top element is the rightmost greater element
        if (!st.empty()) {
            res[i] = nums[st.top()];
            cout << "  Rightmost greater element for " << nums[i] << " is " << res[i] << endl;
        } else {
            cout << "  No greater element found for " << nums[i] << endl;
        }

        // Push current element index to stack
        st.push(i);
        cout << "  Pushing " << nums[i] << " to stack\n";
    }

    return res;
}

int main() {
    int n;
    cout << "Enter number of elements: ";
    cin >> n;
    vector<int> nums(n);

    cout << "Enter elements: ";
    for (int i = 0; i < n; i++) {
        cin >> nums[i];
    }

    vector<int> result = rightMostGreaterElement(nums);

    cout << "\nFinal result: ";
    for (int num : result) {
        cout << num << " ";
    }
    cout << endl;

    return 0;
}

