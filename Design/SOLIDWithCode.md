### **SOLID Principles in C++ (Using a Coffee Maker Example)**
The **SOLID** principles are a set of best practices to design maintainable, scalable, and flexible software.

---

## **1. Single Responsibility Principle (SRP)**
**Definition**: A class should have only one reason to change, meaning it should have **only one responsibility**.

### **Bad Example (Violating SRP)**
```cpp
class CoffeeMaker {
public:
    void brewCoffee() {
        std::cout << "Brewing coffee..." << std::endl;
    }

    void displayStatus() {
        std::cout << "Coffee is ready!" << std::endl;
    }

    void logActivity() {
        std::cout << "Logging: Coffee was made." << std::endl;
    }
};
```
👉 **Problem**: The class **does too many things**—brewing coffee, displaying status, and logging.

### **Good Example (Following SRP)**
```cpp
class CoffeeBrewer {
public:
    void brewCoffee() {
        std::cout << "Brewing coffee..." << std::endl;
    }
};

class CoffeeDisplay {
public:
    void showStatus() {
        std::cout << "Coffee is ready!" << std::endl;
    }
};

class CoffeeLogger {
public:
    void logActivity() {
        std::cout << "Logging: Coffee was made." << std::endl;
    }
};
```
✅ **Now each class has a single responsibility**!

---

## **2. Open/Closed Principle (OCP)**
**Definition**: A class should be open for extension but closed for modification.

### **Bad Example (Violating OCP)**
```cpp
class CoffeeMaker {
public:
    void brewCoffee(const std::string& type) {
        if (type == "Espresso") {
            std::cout << "Brewing Espresso" << std::endl;
        } else if (type == "Latte") {
            std::cout << "Brewing Latte" << std::endl;
        } else {
            std::cout << "Unknown coffee type!" << std::endl;
        }
    }
};
```
👉 **Problem**: Every time a new coffee type is added, we must **modify** the class.

### **Good Example (Following OCP using Inheritance & Polymorphism)**
```cpp
class Coffee {
public:
    virtual void brew() const = 0;  // Pure virtual function
    virtual ~Coffee() = default;
};

class Espresso : public Coffee {
public:
    void brew() const override {
        std::cout << "Brewing Espresso" << std::endl;
    }
};

class Latte : public Coffee {
public:
    void brew() const override {
        std::cout << "Brewing Latte" << std::endl;
    }
};

// Coffee Maker remains unchanged
class CoffeeMaker {
public:
    void brewCoffee(const Coffee& coffee) {
        coffee.brew();
    }
};
```
✅ **Now, we can add new coffee types without modifying `CoffeeMaker`!**

---

## **3. Liskov Substitution Principle (LSP)**
**Definition**: Derived classes should be **substitutable** for their base class without breaking the program.

### **Bad Example (Violating LSP)**
```cpp
class Coffee {
public:
    virtual void brew() const {
        std::cout << "Brewing generic coffee" << std::endl;
    }
};

class InstantCoffee : public Coffee {
public:
    void brew() const override {
        throw std::logic_error("Instant coffee doesn't need brewing!");
    }
};
```
👉 **Problem**: The subclass (`InstantCoffee`) **breaks** the contract of the base class. If `brew()` is called, it **throws an exception**.

### **Good Example (Following LSP)**
```cpp
class Coffee {
public:
    virtual void prepare() const = 0;  // Abstract base class
    virtual ~Coffee() = default;
};

class Espresso : public Coffee {
public:
    void prepare() const override {
        std::cout << "Brewing Espresso" << std::endl;
    }
};

class InstantCoffee : public Coffee {
public:
    void prepare() const override {
        std::cout << "Mixing Instant Coffee with hot water" << std::endl;
    }
};

// Now all subclasses behave correctly
void makeCoffee(const Coffee& coffee) {
    coffee.prepare();
}
```
✅ **Now, all subclasses behave correctly when used as a `Coffee` object.**

---

## **4. Interface Segregation Principle (ISP)**
**Definition**: A class should **not** be forced to implement interfaces it does not use.

### **Bad Example (Violating ISP)**
```cpp
class CoffeeMachine {
public:
    virtual void brewCoffee() = 0;
    virtual void addMilk() = 0;
    virtual void frothMilk() = 0;
};
```
👉 **Problem**: Not all coffee machines have a milk frother! **Espresso machines don’t need `frothMilk()`**.

### **Good Example (Following ISP using Multiple Interfaces)**
```cpp
class CoffeeBrewer {
public:
    virtual void brewCoffee() = 0;
    virtual ~CoffeeBrewer() = default;
};

class MilkFrother {
public:
    virtual void frothMilk() = 0;
    virtual ~MilkFrother() = default;
};

class EspressoMachine : public CoffeeBrewer {
public:
    void brewCoffee() override {
        std::cout << "Brewing Espresso" << std::endl;
    }
};

class CappuccinoMachine : public CoffeeBrewer, public MilkFrother {
public:
    void brewCoffee() override {
        std::cout << "Brewing Cappuccino" << std::endl;
    }
    void frothMilk() override {
        std::cout << "Frothing Milk for Cappuccino" << std::endl;
    }
};
```
✅ **Now, machines only implement the interfaces they need!**

---

## **5. Dependency Inversion Principle (DIP)**
**Definition**: High-level modules should **not** depend on low-level modules. Instead, both should depend on abstractions.

### **Bad Example (Violating DIP)**
```cpp
class EspressoMachine {
public:
    void brewCoffee() {
        std::cout << "Brewing Espresso" << std::endl;
    }
};

class CoffeeShop {
    EspressoMachine machine;  // Direct dependency on a concrete class
public:
    void serveCoffee() {
        machine.brewCoffee();
    }
};
```
👉 **Problem**: `CoffeeShop` is **tightly coupled** to `EspressoMachine`. We can't easily switch to another coffee type.

### **Good Example (Following DIP using Dependency Injection)**
```cpp
class CoffeeBrewer {
public:
    virtual void brewCoffee() = 0;
    virtual ~CoffeeBrewer() = default;
};

class EspressoMachine : public CoffeeBrewer {
public:
    void brewCoffee() override {
        std::cout << "Brewing Espresso" << std::endl;
    }
};

class CoffeeShop {
    std::shared_ptr<CoffeeBrewer> machine;  // Depend on an abstraction
public:
    CoffeeShop(std::shared_ptr<CoffeeBrewer> brewer) : machine(std::move(brewer)) {}
    void serveCoffee() {
        machine->brewCoffee();
    }
};

int main() {
    std::shared_ptr<CoffeeBrewer> espressoMachine = std::make_shared<EspressoMachine>();
    CoffeeShop shop(espressoMachine);
    shop.serveCoffee();
}
```
✅ **Now `CoffeeShop` can work with any `CoffeeBrewer`, making it flexible and testable!**

---

## **Conclusion**
| **SOLID Principle** | **Key Idea** | **Applied in Coffee Maker Example** |
|---------------------|-------------|--------------------------------------|
| **S**ingle Responsibility | One class = One responsibility | `CoffeeBrewer`, `CoffeeDisplay`, `CoffeeLogger` |
| **O**pen/Closed | Open for extension, closed for modification | `Coffee` base class with derived classes |
| **L**iskov Substitution | Derived classes should work as base classes | `InstantCoffee` properly implementing `Coffee` |
| **I**nterface Segregation | Clients should not be forced to implement unnecessary methods | `CoffeeBrewer` & `MilkFrother` interfaces |
| **D**ependency Inversion | Depend on abstractions, not concrete classes | `CoffeeBrewer` base class with `CoffeeShop` |

---

🔥 **By applying SOLID principles, our Coffee Maker is now scalable, maintainable, and testable!** 🚀

Would you like a **Factory Pattern** implementation on top of this for Coffee Creation? 😃
