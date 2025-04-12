#include <iostream>
#include <mutex>

class Singleton {
  public:
    static Singleton& getInstance() {
      std::call_once(initFlag, []() {
          instance = new Singleton();
          });
      return *instance;
    }

    void show() {
      std::cout << "Singleton instance address: " << this << std::endl;
    }
  private:
    Singleton() {}
    ~Singleton() {}
    Singleton(const Singleton&) = delete;
    Singleton& operator=(const Singleton&) = delete;
    static std::once_flag initFlag;
    static Singleton* instance;
    static int count;
};

std::once_flag Singleton::initFlag;
Singleton* Singleton::instance = nullptr;

int main() {
    Singleton& s1 = Singleton::getInstance();
    Singleton& s2 = Singleton::getInstance();

    s1.show();
    s2.show();  // Both should print the same address

    return 0;
}
