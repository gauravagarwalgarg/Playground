#include <iostream>
#include <mutex>

class Singleton {
  public:
    static Singleton* getInstance() {
      std::lock_guard<std::mutex> lock(mtx);
      if(!instance) {
        instance = new Singleton();
      }

      return instance;
    }

    void show() {
      std::cout << "Singleton instance address" << this << std::endl;
    }
  private:
    Singleton() {}
    ~Singleton() {}

    Singleton(const Singleton&) = delete;
    Singleton& operator=(const Singleton&) = delete;

    static Singleton* instance;
    static std::mutex mtx;
};

Singleton* Singleton::instance = nullptr;
std::mutex Singleton::mtx;

int main() {
  Singleton* s1 = Singleton::getInstance();
  Singleton* s2 = Singleton::getInstance();

  s1->show();
  s2->show();

  return 0;
}

