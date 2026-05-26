#include <iostream>

template<typename T>
class UniquePtr {
  private:
    T* ptr;
  public:
    explicit UniquePtr(T* p = nullptr) : ptr(p) {}

    ~UniquePtr() {
      delete ptr;
    }

    UniquePtr(const UniquePtr&) = delete;
    UniquePtr& operator=(const UniquePtr&) = delete;

    UniquePtr(UniquePtr&& other) noexcept : ptr(other.ptr) {
      other.ptr = nullptr;
    }

    UniquePtr& operator=(UniquePtr&& other) noexcept {
      if(this != &other) {
        delete ptr;
        ptr = other.ptr;
        other.ptr = nullptr;
      }
      return *this;
    }

    T& operator*() const { return *ptr; }
    T* operator->() const { return ptr; }

    T* get() const { return ptr; }

    T* release() {
      T* temp = ptr;
      ptr = nullptr;
      return temp;
    }

    void reset(T* p = nullptr) {
      delete ptr;
      ptr = p;
    }
};


