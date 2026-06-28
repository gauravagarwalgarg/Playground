#include <iostream>

// UniquePtr class definition
template <typename T>
class UniquePtr {
private:
    T* ptr;  // Raw pointer

public:
    // Constructor
    explicit UniquePtr(T* p = nullptr) : ptr(p) {}

    // Destructor (deletes managed object)
    ~UniquePtr() {
        delete ptr;
    }

    // Delete copy constructor (No copying allowed)
    UniquePtr(const UniquePtr&) = delete;

    // Delete copy assignment operator
    UniquePtr& operator=(const UniquePtr&) = delete;

    // Move constructor (Transfer ownership)
    UniquePtr(UniquePtr&& other) noexcept : ptr(other.ptr) {
        other.ptr = nullptr;  // Set other's pointer to null
    }

    // Move assignment operator
    UniquePtr& operator=(UniquePtr&& other) noexcept {
        if (this != &other) {
            delete ptr;       // Delete existing resource
            ptr = other.ptr;  // Transfer ownership
            other.ptr = nullptr;
        }
        return *this;
    }

    // Overload * operator (dereference)
    T& operator*() const {
        return *ptr;
    }

    // Overload -> operator (member access)
    T* operator->() const {
        return ptr;
    }

    // Get raw pointer
    T* get() const {
        return ptr;
    }

    // Release ownership and return raw pointer
    T* release() {
        T* temp = ptr;
        ptr = nullptr;
        return temp;
    }

    // Reset with new pointer
    void reset(T* p = nullptr) {
        delete ptr;
        ptr = p;
    }
};


