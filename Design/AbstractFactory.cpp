#include <iostream>
using namespace std;

// Define Abstract products
class Button {
  public:
    virtual void render() = 0;
    virtual ~Button() {}
};

class Checkbox {
  public:
    virtual void render() = 0;
    virtual ~Checkbox() {}
};

// Define Concrete Products
class WindowsButton : public Button {
  public:
    void render() override {
      cout << "Rendering Windows Button" << endl;
    }
};

class MacButton : public Button {
  public:
    void render() override {
      cout << "Rendering MAC button" << endl;
    }
};

class WindowsCheckbox : public Checkbox {
  public:
    void render() override {
      cout << "Rendering Windows checkbox " << endl;
    }
};

class MacCheckBox : public Checkbox {
  public:
    void render() override {
      cout << "Rendering MAC Checkbox" << endl;
    }
};

class GUIFactory {
  public:
    virtual Button* createButton() = 0;
    virtual Checkbox* createCheckbox() = 0;
    virtual ~GUIFactory() {}
};

class WindowsFactory : public GUIFactory {
  public:
    Button* createButton() override {
      return new WindowsButton();
    }

    Checkbox* createCheckbox() override {
      return new WindowsCheckbox();
    }
};

class MacFactory : public GUIFactory {
  public:
    Button* createButton() override {
      return new MacButton();
    }

    Checkbox* createCheckbox() override {
      return new MacCheckBox();
    }
};

class Application {
  private:
    GUIFactory* factory;
  public:
    Application(GUIFactory* f) : factory(f) {}

    void renderUI() {
      Button* btn = factory->createButton();
      Checkbox* cb = factory->createCheckbox();

      btn->render();
      cb->render();

      delete btn;
      delete cb;
    }
};

int main() {
  GUIFactory* factory = nullptr;

  factory = new WindowsFactory();
  Application app1(factory);
  app1.renderUI();
  delete factory;

  factory = new MacFactory();
  Application app2(factory);
  app2.renderUI();
  delete factory;

  return 0;
}

