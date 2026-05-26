"""
Observer Design Pattern
Defines a one-to-many dependency between objects so that when one object
changes state, all its dependents are notified.
"""

from abc import ABC, abstractmethod


class Observer(ABC):
    @abstractmethod
    def update(self, event: str, data: any) -> None:
        pass


class EventManager:
    """Subject / Publisher"""

    def __init__(self):
        self._listeners: dict[str, list[Observer]] = {}

    def subscribe(self, event: str, listener: Observer) -> None:
        if event not in self._listeners:
            self._listeners[event] = []
        self._listeners[event].append(listener)

    def unsubscribe(self, event: str, listener: Observer) -> None:
        if event in self._listeners:
            self._listeners[event].remove(listener)

    def notify(self, event: str, data: any) -> None:
        for listener in self._listeners.get(event, []):
            listener.update(event, data)


class EmailNotifier(Observer):
    def __init__(self):
        self.messages = []

    def update(self, event: str, data: any) -> None:
        self.messages.append(f"Email: [{event}] {data}")


class LogNotifier(Observer):
    def __init__(self):
        self.logs = []

    def update(self, event: str, data: any) -> None:
        self.logs.append(f"Log: [{event}] {data}")


if __name__ == "__main__":
    manager = EventManager()
    email = EmailNotifier()
    logger = LogNotifier()

    manager.subscribe("order_placed", email)
    manager.subscribe("order_placed", logger)
    manager.subscribe("order_shipped", email)

    manager.notify("order_placed", "Order #123")
    manager.notify("order_shipped", "Order #123")

    assert len(email.messages) == 2
    assert email.messages[0] == "Email: [order_placed] Order #123"
    assert email.messages[1] == "Email: [order_shipped] Order #123"
    assert len(logger.logs) == 1
    assert logger.logs[0] == "Log: [order_placed] Order #123"

    # Test unsubscribe
    manager.unsubscribe("order_placed", email)
    manager.notify("order_placed", "Order #456")
    assert len(email.messages) == 2  # no new message
    assert len(logger.logs) == 2

    print("All tests passed!")
