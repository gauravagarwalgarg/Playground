#include <iostream>
#include <vector>
#include <queue>

using namespace std;

void printHeap(std::priority_queue<int> q) {

	std::vector<int> nums;
	while(!q.empty())
	{
		nums.push_back(q.top());
		q.pop();
	}

	std::cout << "Heap state: ";
	for (int num : nums)
	{
		std::cout << num << " ";
	}
}

std::vector<int> findKSmallest(const std::vector<int>& arr, int k) {
	std::priority_queue<int> maxHeap;

	std::cout << "Processing elements: \n";

	for (int num : arr) {
		maxHeap.push(num);
		std::cout << "inserted " << num << " into heap.\n";
		printHeap(maxHeap); // Debug print heap state after removal

		if(maxHeap.size() > k) {
			std::cout << "Removing " << maxHeap.top() << "(smallest in heap) to maintain size" << k << ".\n";
			maxHeap.pop();
			printHeap(maxHeap); // Debug print heap state after removal


		}
	}
	std::vector<int> result;
	while(!maxHeap.empty())
	{
		result.push_back(maxHeap.top());
		maxHeap.pop();
	}
	return result;
}

int main() {
	std::vector<int> nums = {5, 1, 9, 3, 7, 6, 9, 2};
	int k = 3;
	std::vector<int> largestNums = findKSmallest(nums, k);

	std::cout << k << " Smallest Numbers: ";
	for (int num : largestNums) {
		std::cout << num << " ";
	}
	std::cout << std::endl;
	return 0;
}
