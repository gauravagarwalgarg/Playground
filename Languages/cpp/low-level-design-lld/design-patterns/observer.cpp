/**
 * Design Pattern: Observer
 *
 * Define a one-to-many dependency between objects so that when one
 * object changes state, all dependents are notified.
 */
#include <iostream>
#include <vector>
#include <string>
#include <algorithm>
#include <cassert>
using namespace std;

class IObserver {
public:
    virtual ~IObserver() = default;
    virtual void update(const string& event) = 0;
};

class ISubject {
public:
    virtual ~ISubject() = default;
    virtual void attach(IObserver* observer) = 0;
    virtual void detach(IObserver* observer) = 0;
    virtual void notify(const string& event) = 0;
};

class StockTicker : public ISubject {
public:
    void attach(IObserver* observer) override {
        observers_.push_back(observer);
    }
    void detach(IObserver* observer) override {
        observers_.erase(remove(observers_.begin(), observers_.end(), observer), observers_.end());
    }
    void notify(const string& event) override {
        for (auto* obs : observers_) obs->update(event);
    }
    void priceChanged(double newPrice) {
        price_ = newPrice;
        notify("Price changed to " + to_string(newPrice));
    }
private:
    vector<IObserver*> observers_;
    double price_ = 0.0;
};

class AlertSystem : public IObserver {
public:
    void update(const string& event) override {
        lastEvent_ = event;
        alertCount_++;
    }
    int alertCount() const { return alertCount_; }
    string lastEvent() const { return lastEvent_; }
private:
    int alertCount_ = 0;
    string lastEvent_;
};

int main() {
    StockTicker ticker;
    AlertSystem alert1, alert2;

    ticker.attach(&alert1);
    ticker.attach(&alert2);
    ticker.priceChanged(150.5);

    assert(alert1.alertCount() == 1);
    assert(alert2.alertCount() == 1);

    ticker.detach(&alert2);
    ticker.priceChanged(155.0);

    assert(alert1.alertCount() == 2);
    assert(alert2.alertCount() == 1); // detached, no update

    cout << "All tests passed!" << endl;
    return 0;
}
