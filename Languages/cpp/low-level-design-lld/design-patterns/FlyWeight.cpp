// The Flyweight pattern is a structural design pattern that aims to minimize memory usage and improve performance by sharing as much data as possible. Instead of creating a large number of similar objects, it reuses existing objects with shared states.
//

#include <iostream>
#include <unordered_map>
#include <memory>

using namespace std;

class Character {
  private:
    char symbol;
    string font;

  public:
    Character(char sym, string f) : symbol(sym), font(f) {}

    void display(int size, string color) {
      cout << "Character: " << symbol
           << " | Font: " << font
           << " | Size: " << size
           << " | Color: " << color << endl;
    }
};

class CharacterFactory {
  private:
    unordered_map<char, shared_ptr<Character>> characters;
  public:
    shared_ptr<Character> getCharacter(char sym, string font) {
      if(characters.find(sym) == characters.end()) {
        characters[sym] = make_shared<Character>(sym, font);
      }
      return characters[sym];
    }
};

int main() {
  CharacterFactory factory;

  auto c1 = factory.getCharacter('A', "Aerial");
  auto c2 = factory.getCharacter('B', "Aerial");
  auto c3 = factory.getCharacter('B', "Aerial");

  c1->display(12, "Red");
  c2->display(14, "Blue");
  c3->display(16, "Green");

  return 0;
}

