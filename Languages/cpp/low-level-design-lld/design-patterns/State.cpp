#include <iostream>
#include <memory>

using namespace std;

// Forward Declaration
class TrafficLightContext;

// Step 1: State Interface
class TrafficLightState {
public:
    virtual void changeLight(TrafficLightContext& context) = 0;  // Pure virtual function
    virtual string getStateName() = 0;
    virtual ~TrafficLightState() = default;
};

// Step 2: Context (Maintains the current state)
class TrafficLightContext {
private:
    shared_ptr<TrafficLightState> currentState;

public:
    TrafficLightContext(shared_ptr<TrafficLightState> initialState) : currentState(initialState) {}

    void setState(shared_ptr<TrafficLightState> state) {
        currentState = state;
    }

    void requestChange() {
        currentState->changeLight(*this);
    }

    void showState() {
        cout << "Current Light: " << currentState->getStateName() << endl;
    }
};

// Step 3: Concrete State Classes
class RedLight : public TrafficLightState {
public:
    void changeLight(TrafficLightContext& context) override;
    string getStateName() override { return "Red"; }
};

class GreenLight : public TrafficLightState {
public:
    void changeLight(TrafficLightContext& context) override;
    string getStateName() override { return "Green"; }
};

class YellowLight : public TrafficLightState {
public:
    void changeLight(TrafficLightContext& context) override;
    string getStateName() override { return "Yellow"; }
};

// Step 4: Implementing State Transitions
void RedLight::changeLight(TrafficLightContext& context) {
    cout << "Changing from Red to Green..." << endl;
    context.setState(make_shared<GreenLight>());
}

void GreenLight::changeLight(TrafficLightContext& context) {
    cout << "Changing from Green to Yellow..." << endl;
    context.setState(make_shared<YellowLight>());
}

void YellowLight::changeLight(TrafficLightContext& context) {
    cout << "Changing from Yellow to Red..." << endl;
    context.setState(make_shared<RedLight>());
}

// Step 5: Client Code
int main() {
    shared_ptr<TrafficLightState> initial = make_shared<RedLight>();
    TrafficLightContext trafficLight(initial);

    for (int i = 0; i < 5; i++) {
        trafficLight.showState();
        trafficLight.requestChange();
    }

    return 0;
}

