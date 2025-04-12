#include <stack>
#include <vector>
#include <iostream>

using namespace std;

vector<int> rightMostGreaterElement(vector<int>& nums) {
  int n = nums.size();
  vector<int> res(n, -1);
  stack<int> st;

  for(int i = n - 1; i >= 0; i--) {
    while(!st.empty() && nums[st.top()] <= nums[i]) {
      st.pop();
    }

    if(!st.empty()) {
      res[i] = nums[st.top()];
    }

    st.push(i);

  }

  return res;
}

int main() {
  vector<int> arr = {5, 18, 17, 16, 21, 2, 4, 3};

  vector<int> result = rightMostGreaterElement(arr);

  for(int num : result) {
    cout << num << ", ";

  }
}

