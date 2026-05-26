#include <iostream>
#include <thread>
#include <mutex>

std::mutex mtx;

void sharedResourceAccess()
{
	std::lock_gurard<std::mutex> lock(mtx);
}

int main() {
	std::thread t1(sharedResourceAccess);
	std::thread t2(sharedResourceAccess);
	t1.join();
	t2.join();

	return 0;
}
