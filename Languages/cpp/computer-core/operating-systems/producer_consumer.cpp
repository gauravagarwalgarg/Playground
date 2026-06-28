/**
 * OS Concept: Producer-Consumer Problem
 *
 * Classic synchronization problem using mutex + condition variables.
 * Demonstrates thread coordination with a bounded buffer.
 */
#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <queue>
#include <cassert>
using namespace std;

class BoundedBuffer {
public:
    explicit BoundedBuffer(int capacity) : capacity_(capacity) {}

    void produce(int item) {
        unique_lock<mutex> lock(mtx_);
        notFull_.wait(lock, [this] { return (int)buffer_.size() < capacity_; });
        buffer_.push(item);
        produced_++;
        notEmpty_.notify_one();
    }

    int consume() {
        unique_lock<mutex> lock(mtx_);
        notEmpty_.wait(lock, [this] { return !buffer_.empty(); });
        int item = buffer_.front();
        buffer_.pop();
        consumed_++;
        notFull_.notify_one();
        return item;
    }

    int produced() const { return produced_; }
    int consumed() const { return consumed_; }

private:
    queue<int> buffer_;
    int capacity_;
    mutex mtx_;
    condition_variable notFull_, notEmpty_;
    int produced_ = 0, consumed_ = 0;
};

int main() {
    BoundedBuffer buffer(5);
    const int NUM_ITEMS = 20;

    thread producer([&] {
        for (int i = 0; i < NUM_ITEMS; i++) {
            buffer.produce(i);
        }
    });

    thread consumer([&] {
        for (int i = 0; i < NUM_ITEMS; i++) {
            int val = buffer.consume();
            assert(val == i);
        }
    });

    producer.join();
    consumer.join();

    assert(buffer.produced() == NUM_ITEMS);
    assert(buffer.consumed() == NUM_ITEMS);

    cout << "All tests passed! Produced and consumed " << NUM_ITEMS << " items." << endl;
    return 0;
}
