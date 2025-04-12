// The Bridge pattern is a structural design pattern that decouples abstraction from implementation so that both can evolve independently. Instead of having tightly coupled code, the Bridge pattern creates a layer of abstraction that allows different implementations to be used interchangeably.
// Purpose: Decouples abstraction from implementation
// Flexiblity: High
// Best for : When there are multiple variants of both abstraction and implementations
//

#include <iostream>
#include <memory>
using namespace std;


// Implementation Interface Device
class Device {
  public:
    virtual void turnOn() = 0;
    virtual void turnOff() = 0;
    virtual void setVolume(int volume) = 0;
    virtual ~Device() {}
};

// Concrete implementation: TV
class TV : public Device {
  public:
    void turnOn() override {
      cout << "TV is now ON" << endl;
    }

    void turnOff() override {
      cout << "TV is now OFF" << endl;
    }

    void setVolume(int volume) override {
      cout << "TV volume set to " << volume << endl;
    }
};

class Radio : public Device {
  public:
    void turnOn() override {
      cout << "Radio is now ON" << endl;
    }

    void turnOff() override {
      cout << "Radio is now OFF" << endl;
    }

    void setVolume(int volume) override {
      cout << "Radio volume set to " << volume << endl;
    }
};

class RemoteControl {
  protected:
    unique_ptr<Device> device;

  public:
    RemoteControl(unique_ptr<Device> dev) : device(move(dev)) {}

    virtual void powerButton() {
      cout << "Power button pressed." << endl;
      device->turnOn();
    }

    virtual void volumeUp() {
      cout << "Increasing volume..." << endl;
      device->setVolume(10);
      }

    virtual ~RemoteControl() {}
};

// Extended Abstraction
class AdvancedRemote: public RemoteControl {
  public:
    AdvancedRemote(unique_ptr<Device> dev) : RemoteControl(move(dev)) {}

    void mute() {
      cout << "Muting device..." << endl;
      device->setVolume(0);
    }
};

int main() {
  unique_ptr<RemoteControl> tvRemote = make_unique<RemoteControl>(make_unique<TV>());
  tvRemote->powerButton();
  tvRemote->volumeUp();

  unique_ptr<AdvancedRemote> radioRemote = make_unique<AdvancedRemote>(make_unique<Radio>());
  radioRemote->powerButton();
  radioRemote->mute();

  return 0;
}
