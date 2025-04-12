#include <iostream>
#include <thread>
#include <future>

void computeSum(std::promise<int> sumPromise, int a, int b) {
    int result = a + b;
    std::this_thread::sleep_for(std::chrono::seconds(1)); // Simulate delay
    sumPromise.set_value(result); // Set the result in promise
}

int main() {
    std::promise<int> sumPromise;
    std::future<int> sumFuture = sumPromise.get_future(); // Get the associated future

    std::thread worker(computeSum, std::move(sumPromise), 10, 20); // Pass the promise to the thread

    std::cout << "Waiting for the result..." << std::endl;
    int result = sumFuture.get(); // Waits until the value is set
    std::cout << "Result: " << result << std::endl;

    worker.join(); // Clean up the thread
    return 0;
}

