#include <iostream>
#include <unordered_map>
#include <vector>
#include <queue>
#include <mutex>
#include <condition_variable>
#include <thread>
#include <chrono>
#include <atomic>

class Message {
public:
    int offset;
    std::string content;
    Message(int off, std::string msg) : offset(off), content(std::move(msg)) {}
};

class Topic {
private:
    std::string name;
    std::queue<Message> messages;
    std::unordered_map<int, int> consumerOffsets;
    std::mutex mtx;
    std::condition_variable cv;
    int offset = 0;
    int retentionPeriod; // in seconds

public:
    explicit Topic(std::string topicName, int retention) : name(std::move(topicName)), retentionPeriod(retention) {}
    
    void publish(const std::string& message) {
        std::lock_guard<std::mutex> lock(mtx);
        messages.emplace(offset++, message);
        cv.notify_all();
    }
    
    std::vector<Message> consume(int consumerId) {
        std::unique_lock<std::mutex> lock(mtx);
        cv.wait(lock, [&] { return !messages.empty(); });
        
        std::vector<Message> result;
        while (!messages.empty()) {
            Message msg = messages.front();
            if (consumerOffsets[consumerId] < msg.offset) {
                result.push_back(msg);
                consumerOffsets[consumerId] = msg.offset;
            }
            messages.pop();
        }
        return result;
    }
    
    void resetOffset(int consumerId, int newOffset) {
        std::lock_guard<std::mutex> lock(mtx);
        consumerOffsets[consumerId] = newOffset;
    }
    
    void pruneOldMessages() {
        std::lock_guard<std::mutex> lock(mtx);
        while (!messages.empty() && messages.front().offset < offset - retentionPeriod) {
            messages.pop();
        }
    }
};

class PubSubSystem {
private:
    std::unordered_map<std::string, std::shared_ptr<Topic>> topics;
    std::mutex topicMtx;

public:
    void createTopic(const std::string& topicName, int retention) {
        std::lock_guard<std::mutex> lock(topicMtx);
        topics[topicName] = std::make_shared<Topic>(topicName, retention);
    }
    
    void deleteTopic(const std::string& topicName) {
        std::lock_guard<std::mutex> lock(topicMtx);
        topics.erase(topicName);
    }
    
    void publishMessage(const std::string& topicName, const std::string& message) {
        std::lock_guard<std::mutex> lock(topicMtx);
        if (topics.find(topicName) != topics.end()) {
            std::thread([&, topicName, message] {
                topics[topicName]->publish(message);
            }).detach();
        }
    }
    
    std::vector<Message> consumeMessages(const std::string& topicName, int consumerId) {
        std::lock_guard<std::mutex> lock(topicMtx);
        if (topics.find(topicName) != topics.end()) {
            return topics[topicName]->consume(consumerId);
        }
        return {};
    }
    
    void resetConsumerOffset(const std::string& topicName, int consumerId, int offset) {
        std::lock_guard<std::mutex> lock(topicMtx);
        if (topics.find(topicName) != topics.end()) {
            topics[topicName]->resetOffset(consumerId, offset);
        }
    }
};

int main() {
    PubSubSystem pubSub;
    pubSub.createTopic("news", 10);
    
    pubSub.publishMessage("news", "Breaking News 1");
    pubSub.publishMessage("news", "Breaking News 2");
    
    std::vector<Message> messages = pubSub.consumeMessages("news", 1);
    for (const auto& msg : messages) {
        std::cout << "Consumer 1 received: " << msg.content << " at offset " << msg.offset << std::endl;
    }
    return 0;
}

