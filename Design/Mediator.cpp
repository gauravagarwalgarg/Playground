#include <iostream>
#include <vector>
#include <memory>
#include <string>

using namespace std;

// Forward Declaration
class User;
class ChatRoom;

// Step 1: Mediator Interface
class Mediator {
public:
    virtual void showMessage(shared_ptr<User> user, const string& message) = 0;
};

// Step 2: Concrete Mediator (Manages communication)
class ChatRoom : public Mediator {
public:
    void showMessage(shared_ptr<User> user, const string& message) override;
};

// Step 3: Colleague (User class that interacts via Mediator)
class User : public enable_shared_from_this<User> {  // FIX: Enable shared_from_this
private:
    string name;
    shared_ptr<Mediator> chatRoom;  // Reference to the Mediator

public:
    User(string n, shared_ptr<Mediator> mediator) : name(n), chatRoom(mediator) {}

    void sendMessage(const string& message) {
        chatRoom->showMessage(shared_from_this(), message);  // FIX: shared_from_this() works now
    }

    string getName() {
        return name;
    }
};

// Define `showMessage` outside class (Circular Dependency)
void ChatRoom::showMessage(shared_ptr<User> user, const string& message) {
    cout << "[" << user->getName() << "]: " << message << endl;
}

// Step 4: Client Code
int main() {
    shared_ptr<Mediator> chatRoom = make_shared<ChatRoom>();

    shared_ptr<User> alice = make_shared<User>("Alice", chatRoom);
    shared_ptr<User> bob = make_shared<User>("Bob", chatRoom);

    alice->sendMessage("Hi Bob!");
    bob->sendMessage("Hello Alice!");

    return 0;
}

