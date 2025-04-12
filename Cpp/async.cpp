#include <iostream>
#include <future>
#include <thread>
#include <chrono>

int multiply(int a, int b) {
  std::this_thread::sleep_for(std::chrono::milliseconds(500));
  return a*b;
}

int main() {
  std::future<int> resultFuture = std::async(std::launch::async, multiply, 5, 6);

  std::cout << "Doing other stuff while waiting...\n";

  int result = resultFuture.get();
  std::cout << "Result: " << result << std::endl;

  return 0;
}
