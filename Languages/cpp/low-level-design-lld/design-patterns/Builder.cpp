#include <iostream>
#include <string>

using namespace std;

class Burger {
  private:
    string bun;
    string patty;
    bool cheese;
    bool lettuce;

  public:
    Burger(string bun, string patty, bool cheese, bool lettuce) :
      bun(bun), patty(patty), cheese(cheese), lettuce(lettuce) {}

    void show() {
      cout << "Burger with " << bun << "bun, " << patty << " patty, "
        << (cheese ? "cheese, " : "no cheese, ")
        << (lettuce ? "lettuce." : "no lettuce.") << endl;
    }
};

class BurgerBuilder {
  private:
    string bun = "Regular";
    string patty = "veg";
    bool cheese = false;
    bool lettuce = false;

  public:
    BurgerBuilder& setBun(string bunType) {
      this->bun = bunType;
      return *this;
    }

    BurgerBuilder& setPatty(string pattyType) {
      this->patty = pattyType;
      return *this;
    }

    BurgerBuilder& addCheese() {
      this->cheese = true;
      return *this;

    }

    BurgerBuilder& addLettuce() {
      this->lettuce = true;
      return *this;
    }

    Burger build() {
      return Burger(bun, patty, cheese, lettuce);
    }
};

int main() {
  Burger classicBurger = BurgerBuilder()
                          .setBun("Sesame")
                          .setPatty("Veg")
                          .addCheese()
                          .build();

  Burger viggieBurger = BurgerBuilder()
                          .setPatty("Veggie")
                          .addLettuce()
                          .build();

  classicBurger.show();
  viggieBurger.show();

  return 0;
}
