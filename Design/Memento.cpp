// The Memento pattern is a behavioral design pattern that allows capturing an object's internal state so that it can be restored later, without violating encapsulation.

#include <iostream>
#include <vector>
#include <memory>

using namespace std;

class Memento {
  private:
    string state;
  public:
    Memento(string s) : state(s) {}

    string getState() const {
      return state;
    }
};

// Originator
class TextEditor {
  private:
    string text;

  public:
    void write(const string& newText) {
      text += newText;
    }

    void show() {
      cout << "Current Text: " << text << endl;
    }

    shared_ptr<Memento> save() {
      return make_shared<Memento>(text);
    }

    void restore(shared_ptr<Memento> memento) {
      if(memento) {
        text = memento->getState();
      }
    }
};

// Caretaker
class History {
  private:
    vector<shared_ptr<Memento>> history;

  public:
    void push(shared_ptr<Memento> memento) {
      history.push_back(memento);
    }

    shared_ptr<Memento> pop() {
      if(history.empty()) return nullptr;
      shared_ptr<Memento> memento = history.back();
      history.pop_back();
      return memento;
    }
};

int main() {
  TextEditor editor;
  History history;

  editor.write("Hello, ");
  history.push(editor.save());

  editor.write("World!");
  history.push(editor.save());

  editor.show();

  cout << "Undoing last change..." << endl;

  editor.restore(history.pop());
  editor.show();

  cout << "Undoing again..." << endl;
  editor.restore(history.pop());
  editor.show();

  return 0;
}

