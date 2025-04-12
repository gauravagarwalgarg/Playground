int minFinalSum(vector<int>& nums, int k) {
  priority_queue<int> maxHeap(nums.begin(), nums.end());

  for(int i = 0; i < k; i++) {
    int maxVal = maxHeap.top();
    maxHeap.pop();
    maxHeap.push(ceil(maxVal / 2.0));
  }

  int sum = 0;
  while(!maxHeap.empty()) {
    sum += maxHeap.top();
    maxHeap.pop();
  }

  return sum;
}
