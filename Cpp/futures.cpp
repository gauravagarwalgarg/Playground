#include <iostream>
#include <thread>
#include <future>

void computeSum(std::promise<int> sumPromise, int a, int b) {
  int result = a + b;
  std::this_thread::sleep_for(std::chrono::seconds(1));
  sumPromise.set_value(result);
}

int main() {
  std::promise<int> sumPromise;
  std::future<int> sumFuture = sumPromise.get_future();

  std::thread worker(computeSum, std::move(sumPromise), 10, 20);

  std::cout << "Waiting for the result..." << std::endl;
  int result = sumFuture.get();
  std::cout << "Result: " << result << std::endl;

  worker.join();
  return 0;
}

