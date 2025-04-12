#include <iostream>
#include <vector>
#include <memory>
#include <thread>
#include <mutex>
#include <future>
#include <map>
#include <algorithm>

class Task {
  public:
    virtual void execute() = 0;
    virtual ~Task() = default;
};

class ComputeTask : public Task {
  int data;
  public:
    explicit ComputeTask(int d) : data(d) {}
    void execute() override {
      std::cout << "ComputeTask processing data: " << data * data << std::endl;
    }
};

class NetworkTask : public Task {
  std::string url;
  public:
    explicit NetworkTask(std::string u) : url(std::move(u)) {}
    void execute() override {
      std::cout <<  "NetworkTask fetching from: " << url << std::endl;
    }
};

class TaskManager {
  private:
    std::vector<std::unique_ptr<Task>> tasks;
    std::mutex taskMutex;
  public:
    void addTask(std::unique_ptr<Task> task) {
      std::lock_guard<std::mutex> lock(taskMutex);
      tasks.push_back(std::move(task));
    }

    void processTasks() {
      std::vector<std::future<void>> futures;
      for(auto & task : tasks) {
        futures.push_back(std::async(std::launch::async, [&task]() { task->execute(); }));
      }
      for(auto & f: futures) {
        f.get();
      }
    }
};


void demonstrateCasting() {
  int a = 10;
  double b = static_cast<double>(a);

  const char* str = "Hello";
  char* modifiable = const_cast<char*>(str);

  int num = 42;
  void* ptr = &num;
  int* numPtr = reinterpret_cast<int*>(ptr);

  std::unique_ptr<Task> base = std::make_unique<ComputeTask>(100);
  ComputeTask* derived = dynamic_cast<ComputeTask*>(base.get());

  std::cout << "Static cast result: " << b << "\n";
  std::cout << "Reinterpret cast result: " << *numPtr << "\n";
  if(derived) {
    std::cout << "Dynamic cast successful : ComputeTask derived from Task\n";
  }
}

int main() {
  TaskManager manager;

  manager.addTask(std::make_unique<ComputeTask>(5));
  manager.addTask(std::make_unique<NetworkTask>("https://example.com"));
  manager.addTask(std::make_unique<ComputeTask>(10));

  std::cout << "procesing tasks asynchronously:\n";
  manager.processTasks();

  std::cout << "\nDemonstrating various type casts:\n";
  demonstrateCasting();

  return 0;
}
