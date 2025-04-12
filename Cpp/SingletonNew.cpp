#include <iostream>
#include <mutex>

class Singleton {
  private:
    static Singleton* instance;
    static std::once_flag initFlag;

    Singleton() {
      std::cout << "Singleton Instance created\n";
    }

  public:
    Singleton(const Singleton&) = delete;
    Singleton& operator=(const Singleton&) = delete;

    static Singleton* getInstance() {
      std::call_once(initFlag, []() {
            instance = new Singleton();
          });
      return instance;
    }

    void showMessage() {
      std::cout << "Address of the Singleton " << this << std::endl;
    }
};

Singleton* Singleton::instance = nullptr;
std::once_flag Singleton::initFlag;

int main() {
  Singleton* s1 = Singleton::getInstance();
  Singleton* s2 = Singleton::getInstance();

  s1->showMessage();
  s2->showMessage();

  return 0;
}
