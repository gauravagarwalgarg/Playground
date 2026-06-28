#include <cmath>
#include <compare>
#include <cstdlib>
#include<bits/stdc++.h>
#include <iostream>
#include <bits/stdc++.h>
#include <span>
#include <unordered_map>

using namespace std;

int smallest_pos_number(const vector<int>&arr) {
    int len = arr.size();
    int max_element = 0;

    for(int i = 0; i < len; i++) {
      if(arr[i] > max_element ) {
        max_element = arr[i];
      }
    }

    vector<bool> visited(max_element + 1, false);

    // positive
    for(int num : arr) {
      if(num > 0 && num <= max_element) {
        visited[num] = true;
      }
    }

    for(int i = 1; i <= max_element; i++) {
      if(!visited[i]) {
        return i;
      }
    }


    return max_element + 1; // edge case for max_element 
}

int main() {
  vector<int> nums={2, -3, 4, 1, 1, 7};
  int pos = smallest_pos_number(nums);

  std::cout << pos << std::endl;
  return 0;
}
