#include <iostream>

class HotBeverage {
  public:
    void prepareRecipe() {
      boilWater();
      brew();
      pourInCup();
      if(customerWantsCondiments()) {
        addCondiments();
      }
    }
  protected:
    virtual void brew() = 0;
    virtual void addCondiments() = 0;

    void boilWater() {
      std::cout << "Boiling water...\n";
    }

    void pourInCup() {
      std::cout << "Pouring into cup...\n";
    }

    virtual bool customerWantsCondiments() {
      return true;
    }
};

class Tea : public HotBeverage {
  protected:
    void brew() override {
      std::cout << "Steeping the tea...\n";
    }

    void addCondiments() override {
      std::cout << "Adding lemon...\n";
    }
};

class Coffee : public HotBeverage {
  protected:
    void brew() override {
      std::cout << "Dripping coffee through filter...\n";
    }

    void addCondiments() override {
      std::cout << "Adding sugar and milk...\n";
    }

    bool customerWantsCondiments() override {
      char choice;
      std::cout << "Do you want milk and sugar? (y/n): ";
      std::cin >> choice;
      return (choice == 'y' || choice == 'Y');
    }
};

int main() {
  std::cout << "Making tea...\n";
  Tea myTea;
  myTea.prepareRecipe();

  std::cout << "Making coffee...\n";
  Coffee mycoffee;
  mycoffee.prepareRecipe();

  return 0;
}

