#include <vector>
#include <stack>
#include <iostream>

using namespace std;

vector<int> Checksum(vector<int>& nums) {
  int n = nums.size();
  vector<int> result(n , -1);
  stack<int> st;

  for(int i = 0; i < n; i++) {
    while(!st.empty() && nums[i] > nums[st.top()]) {
      int idx = st.top();
      st.pop();
      result[idx] = nums[i];
    }

    st.push(i);
  }

  return result;
}

int main() {
  vector<int> arr = {5, 18, 16, 17, 21, 2, 3, 4};

  vector<int> result = Checksum(arr);

  for(int num : result) {
    cout << num << ", ";
  }
  cout << "\n";
}

