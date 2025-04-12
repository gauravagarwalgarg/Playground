#include <iostream>

class Singleton {
  public:
    static Singleton& getInstance() {
      static Singleton instance;
      return instance;
    }

    void show() {
      std::cout << "Singleton instance address: " << this << std::endl;
    }

  private:
    Singleton() {}
    ~Singleton() {}
    Singleton(const Singleton&) = delete;
    Singleton& operator=(const Singleton&) = delete;
};

int main() {
  Singleton& s1 = Singleton::getInstance();
  Singleton& s2 = Singleton::getInstance();

  s1.show();
  s2.show();

  return 0;
}
