"""
Mediator Pattern

Defines an object that encapsulates how a set of objects interact.
ChatRoom mediates communication between Users, preventing direct coupling.
"""

from dataclasses import dataclass, field


class ChatRoom:
    """Mediator: routes messages between users."""

    def __init__(self, name: str):
        self.name = name
        self._users: dict[str, "User"] = {}
        self.message_log: list[str] = []

    def join(self, user: "User") -> str:
        self._users[user.name] = user
        user.room = self
        return f"{user.name} joined {self.name}"

    def leave(self, user: "User") -> str:
        self._users.pop(user.name, None)
        user.room = None
        return f"{user.name} left {self.name}"

    def send(self, message: str, sender: "User", recipient: str | None = None) -> None:
        if recipient:
            # Direct message
            if user := self._users.get(recipient):
                user.receive(message, sender.name)
                self.message_log.append(f"{sender.name} -> {recipient}: {message}")
        else:
            # Broadcast
            for name, user in self._users.items():
                if name != sender.name:
                    user.receive(message, sender.name)
            self.message_log.append(f"{sender.name} -> all: {message}")


@dataclass
class User:
    name: str
    room: ChatRoom | None = field(default=None, repr=False)
    inbox: list[str] = field(default_factory=list)

    def send(self, message: str, to: str | None = None) -> None:
        if self.room:
            self.room.send(message, self, to)

    def receive(self, message: str, sender: str) -> None:
        self.inbox.append(f"{sender}: {message}")


if __name__ == "__main__":
    room = ChatRoom("dev-chat")
    alice = User("Alice")
    bob = User("Bob")
    charlie = User("Charlie")

    room.join(alice)
    room.join(bob)
    room.join(charlie)

    # Broadcast
    alice.send("Hello everyone!")
    assert "Alice: Hello everyone!" in bob.inbox
    assert "Alice: Hello everyone!" in charlie.inbox
    assert len(alice.inbox) == 0  # sender doesn't receive own message

    # Direct message
    bob.send("Hey Alice!", to="Alice")
    assert "Bob: Hey Alice!" in alice.inbox
    assert len(charlie.inbox) == 1  # charlie didn't get DM

    # Leave room
    room.leave(charlie)
    alice.send("Charlie left!")
    assert len(charlie.inbox) == 1  # no new messages

    assert len(room.message_log) == 3

    print("All tests passed!")
