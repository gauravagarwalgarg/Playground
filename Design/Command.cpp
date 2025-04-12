#include <iostream>
#include <memory>
#include <vector>

using namespace std;


// Command Interface
class Command {
  public:
    virtual void execute() = 0;
    virtual void undo() = 0;
    virtual ~Command() {}
};

class Light {
  public:
    void turnOn() {
      cout << "Light is ON" << endl;
    }

    void turnOff() {
      cout << "Light is OFF" << endl;
    }

};

class Fan {
  public:
    void increaseSpeed() {
      cout << "Fan Speed increased" << endl;
    }

    void decreaseSpeed() {
      cout << "Fan speed decreased" << endl;
    }
};

class LightOnCommand : public Command {
  private:
    shared_ptr<Light> light;

  public:
    LightOnCommand(shared_ptr<Light> l) : light(l) {}

    void execute() override {
      light->turnOn();
    }

    void undo() override {
      light->turnOff();
    }
};

class LightOffCommand : public Command {
  private:
    shared_ptr<Light> light;

  public:
    LightOffCommand(shared_ptr<Light> l) : light(l) {}

    void execute() override {
      light->turnOff();
    }

    void undo() override {
      light->turnOn();
    }
};


class FanIncreaseCommand : public Command {
  private:
    shared_ptr<Fan> fan;
  public:
    FanIncreaseCommand(shared_ptr<Fan> f) : fan(f) {}

    void execute() override {
      fan->increaseSpeed();
    }

    void undo() override {
      fan->decreaseSpeed();
    }
};

class RemoteControl {
  private:
    vector<shared_ptr<Command>> history;

  public:
    void pressButton(shared_ptr<Command> command) {
      command->execute();
      history.push_back(command);
    }

    void pressUndo() {
      if(!history.empty()) {
        history.back()->undo();
        history.pop_back();
      } else {
        cout << "Nothing to undo!" << endl;
      }
    }
};

int main() {
  shared_ptr<Light> light = make_shared<Light>();
  shared_ptr<Fan> fan = make_shared<Fan>();

  shared_ptr<Command> lightOn = make_shared<LightOnCommand>(light);
  shared_ptr<Command> lightOff = make_shared<LightOffCommand>(light);
  shared_ptr<Command> fanIncrease = make_shared<FanIncreaseCommand>(fan);

  RemoteControl remote;

  remote.pressButton(lightOn);
  remote.pressButton(fanIncrease);
  remote.pressUndo();
  remote.pressUndo();
  return 0;
}


