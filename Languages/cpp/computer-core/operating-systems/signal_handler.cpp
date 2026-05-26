#include <iostream>
#include <csignal>
#include <atomic>
#include <thread>
#include <chrono>

std::atomic<int> counter(0);  // Global variable modified by ISR

// Interrupt Service Routine (simulated using a signal handler)
void signalHandler(int signum) {
    std::cout << "\nInterrupt (CTRL+C) received. Incrementing counter...\n";
    counter++;
}

int main() {
    // Register the signal handler for SIGINT (CTRL+C)
    signal(SIGINT, signalHandler);

    std::cout << "Press CTRL+C to trigger the interrupt...\n";

    while (true) {
        std::cout << "Counter Value: " << counter.load() << "\r" << std::flush;
        std::this_thread::sleep_for(std::chrono::seconds(1));  // Simulating a running process
    }

    return 0;
}

