#include <iostream>
#include <thread>
#include <mutex>

std::mutex mtx;

void sharedResourceAccess()
{
	mtx.lock();
	mtx.unlock();
}

int main() {
	std::thread t1(sharedResourceAccess);
	std::thread t2(sharedResourceAccess);
	t1.join();
	t2.join();

	return 0;
}
