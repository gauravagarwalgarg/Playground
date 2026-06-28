#include <iostream>
#include <vector>
#include <queue>

void printHeap(std::priority_queue<int, std::vector<int>, std::greater<int>> minHeap) {
    std::vector<int> heapElements;
    while (!minHeap.empty()) {
        heapElements.push_back(minHeap.top());
        minHeap.pop();
    }

    std::cout << "Heap state: ";
    for (int num : heapElements) {
        std::cout << num << " ";
    }
    std::cout << std::endl;
}

std::vector<int> findKLargest(const std::vector<int>& arr, int k) {
    std::priority_queue<int, std::vector<int>, std::greater<int>> minHeap;

    std::cout << "Processing elements:\n";
    for (int num : arr) {
        minHeap.push(num);
        std::cout << "Inserted " << num << " into heap.\n";
        printHeap(minHeap); // Debug print heap state

        if (minHeap.size() > k) {
            std::cout << "Removing " << minHeap.top() << " (smallest in heap) to maintain size " << k << ".\n";
            minHeap.pop();
            printHeap(minHeap); // Debug print heap state after removal
        }
    }

    std::vector<int> result;
    while (!minHeap.empty()) {
        result.push_back(minHeap.top());
        minHeap.pop();
    }
    return result;
}

int main() {
    std::vector<int> nums = {5, 1, 9, 3, 7, 6, 9, 2};
    int k = 3;
    std::vector<int> largestNums = findKLargest(nums, k);

    std::cout << k << " Largest Numbers: ";
    for (int num : largestNums) {
        std::cout << num << " ";
    }
    std::cout << std::endl;
    return 0;
}

