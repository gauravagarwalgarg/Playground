#include <iostream>
#include <vector>
#include <memory>
#include <algorithm>

// Observer Interface
class IObserver {
public:
    virtual ~IObserver() = default;
    virtual void update(int data) = 0;
};

// Subject (Observable) Interface
class ISubject {
public:
    virtual ~ISubject() = default;
    virtual void attach(std::shared_ptr<IObserver> observer) = 0;
    virtual void detach(std::shared_ptr<IObserver> observer) = 0;
    virtual void notify() = 0;
};

// Concrete Subject
class ConcreteSubject : public ISubject {
private:
    std::vector<std::shared_ptr<IObserver>> observers;
    int state;
public:
    void attach(std::shared_ptr<IObserver> observer) override {
        observers.push_back(observer);
    }

    void detach(std::shared_ptr<IObserver> observer) override {
        observers.erase(std::remove(observers.begin(), observers.end(), observer), observers.end());
    }

    void notify() override {
        std::cout << "[Subject] Notifying observers...\n";
        for (const auto& observer : observers) {
            observer->update(state);
        }
    }

    void setState(int newState) {
        state = newState;
        std::cout << "[Subject] State changed to: " << state << "\n";
        notify();  // Notify observers of the state change
    }
};

// Concrete Observer
class ConcreteObserver : public IObserver {
private:
    std::string name;
public:
    ConcreteObserver(const std::string& observerName) : name(observerName) {}

    void update(int data) override {
        std::cout << "[Observer " << name << "] Received update: " << data << "\n";
    }
};

int main() {
    auto subject = std::make_shared<ConcreteSubject>();

    auto observer1 = std::make_shared<ConcreteObserver>("A");
    auto observer2 = std::make_shared<ConcreteObserver>("B");
    auto observer3 = std::make_shared<ConcreteObserver>("C");

    subject->attach(observer1);
    subject->attach(observer2);
    subject->attach(observer3);

    std::cout << "\n--- Changing state to 10 ---\n";
    subject->setState(10);

    std::cout << "\n--- Detaching Observer B ---\n";
    subject->detach(observer2);

    std::cout << "\n--- Changing state to 20 ---\n";
    subject->setState(20);

    return 0;
}

