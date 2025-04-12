#include <iostream>
#include <mutex>
#include <vector>

using namespace std;

class FreeListAllocator {
	private:
		struct Node {
			Node* next;
		};

		void* memory;
		Node* freeList;
		size_t blockSize;
		size_t blockCount;
		mutex mtx;

	public:

