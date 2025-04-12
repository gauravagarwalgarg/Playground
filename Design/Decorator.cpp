#include <iostream>
#include <memory>

using namespace std;

class Coffee {
  public:
    virtual string getDescription() const = 0;
    virtual double cost() const = 0;
    virtual ~Coffee() {}
};

class SimpleCoffee : public Coffee {
  public:
    string getDescription() const override {
      return "Simple Coffee";
    }

    double cost() const override {
      return 5.0;
    }
};

class CoffeeDecorator : public Coffee {
  protected:
    unique_ptr<Coffee> coffee;

  public:
    CoffeeDecorator(unique_ptr<Coffee> coffee) :coffee(move(coffee)) {}

    string getDescription() const override {
      return coffee->getDescription();
    }

    double cost() const override {
      return coffee->cost();
    }
};

class Milk : public CoffeeDecorator {
  public:
    Milk(unique_ptr<Coffee> coffee) : CoffeeDecorator(move(coffee)) {}

    string getDescription() const override {
      return coffee->getDescription() + ", Milk";
    }

    double cost() const override {
      return coffee->cost() + 1.5;
    }
};

class Sugar : public CoffeeDecorator {
  public:
    Sugar(unique_ptr<Coffee> coffee) : CoffeeDecorator(move(coffee)) {}

    string getDescription() const override {
      return coffee->getDescription() + ", Sugar";
    }

    double cost() const override {
      return coffee->cost() + 0.5;
    }

};

int main() {
  unique_ptr<Coffee> coffee = make_unique<SimpleCoffee>();
  cout << coffee->getDescription() << " -> Cost: $" << coffee->cost() << endl;

  coffee = make_unique<Milk>(move(coffee));
  cout << coffee->getDescription() << " -> Cost: $" << coffee->cost() << endl;

  coffee = make_unique<Sugar>(move(coffee));
  cout << coffee->getDescription() << " -> Cost: $" << coffee->cost() << endl;

  return 0;
}
