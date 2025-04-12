// The Strategy Pattern is a behavioral design pattern that allows a family of algorithms to be defined and encapsulated in separate classes, enabling them to be interchangeable at runtime.
//

#include <iostream>
#include <memory>

using namespace std;

class PaymentStrategy {
  public:
    virtual void pay(int amount) = 0;
    virtual ~PaymentStrategy() = default;
};

class CreditCardPayment : public PaymentStrategy {
  public:
    void pay(int amount) override {
      cout << "Paid $ " << amount << "using credit card." << endl;
    }
};

class PayPalPayment : public PaymentStrategy {
  public:
    void pay(int amount) override {
      cout << "Paid $ " << amount << "using Paypal." << endl;
    }
};

class BitcoinPayment : public PaymentStrategy {
  public:
    void pay(int amount) override {
      cout << "Paid $ " << amount << "using Bitcoin." << endl;
    }
};

class PaymentContext {
  private:
    shared_ptr<PaymentStrategy> strategy;

  public:
    void setStrategy(shared_ptr<PaymentStrategy> strat) {
      strategy = strat;
    }

  void processPayment(int amount) {
    if(strategy) {
      strategy->pay(amount);
    } else {
      cout << "No payment stratergy selected" << endl;
    }
  }
};

int main() {
  PaymentContext payment;

  payment.setStrategy(make_shared<CreditCardPayment>());
  payment.processPayment(100);


  payment.setStrategy(make_shared<PayPalPayment>());
  payment.processPayment(250);

  payment.setStrategy(make_shared<BitcoinPayment>());
  payment.processPayment(500);

  return 0;
}

