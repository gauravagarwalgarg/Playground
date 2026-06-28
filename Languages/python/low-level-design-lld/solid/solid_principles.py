"""
SOLID Principles - Before/After Examples

Each principle demonstrated with a violation and its fix.
S - Single Responsibility
O - Open/Closed
L - Liskov Substitution
I - Interface Segregation
D - Dependency Inversion
"""

from abc import ABC, abstractmethod
from dataclasses import dataclass
from typing import Protocol
import math


# =============================================================================
# S: Single Responsibility Principle
# A class should have only one reason to change.
# =============================================================================

# ❌ VIOLATION: User handles both auth and profile concerns
class UserMonolith:
    def __init__(self, name: str, email: str, password: str):
        self.name = name
        self.email = email
        self.password = password

    def authenticate(self, password: str) -> bool:
        return self.password == password

    def update_profile(self, name: str) -> None:
        self.name = name


# ✅ FIX: Split into focused classes
@dataclass
class UserProfile:
    name: str
    email: str

    def update_name(self, name: str) -> None:
        self.name = name


class UserAuth:
    def __init__(self, email: str, password_hash: str):
        self.email = email
        self._password_hash = password_hash

    def authenticate(self, password_hash: str) -> bool:
        return self._password_hash == password_hash


# =============================================================================
# O: Open/Closed Principle
# Open for extension, closed for modification.
# =============================================================================

# ❌ VIOLATION: Adding new shapes requires modifying area_calculator
def bad_area_calculator(shapes: list[dict]) -> float:
    total = 0.0
    for s in shapes:
        if s["type"] == "circle":
            total += math.pi * s["radius"] ** 2
        elif s["type"] == "rectangle":
            total += s["width"] * s["height"]
        # Must modify this function for every new shape!
    return total


# ✅ FIX: Polymorphism - new shapes just implement area()
class Shape(ABC):
    @abstractmethod
    def area(self) -> float:
        pass


class Circle(Shape):
    def __init__(self, radius: float):
        self.radius = radius

    def area(self) -> float:
        return math.pi * self.radius ** 2


class Rectangle(Shape):
    def __init__(self, width: float, height: float):
        self.width = width
        self.height = height

    def area(self) -> float:
        return self.width * self.height


def area_calculator(shapes: list[Shape]) -> float:
    return sum(s.area() for s in shapes)


# =============================================================================
# L: Liskov Substitution Principle
# Subtypes must be substitutable for their base types.
# =============================================================================

# ❌ VIOLATION: Square breaks Rectangle's contract
class RectangleBase:
    def __init__(self, width: float, height: float):
        self._width = width
        self._height = height

    @property
    def width(self) -> float:
        return self._width

    @width.setter
    def width(self, value: float) -> None:
        self._width = value

    @property
    def height(self) -> float:
        return self._height

    @height.setter
    def height(self, value: float) -> None:
        self._height = value

    def area(self) -> float:
        return self._width * self._height


class BadSquare(RectangleBase):
    """Violates LSP: setting width also changes height (surprising behavior)."""
    @RectangleBase.width.setter
    def width(self, value: float) -> None:
        self._width = self._height = value

    @RectangleBase.height.setter
    def height(self, value: float) -> None:
        self._width = self._height = value


# ✅ FIX: Separate types, shared interface
class ShapeLSP(ABC):
    @abstractmethod
    def area(self) -> float:
        pass


@dataclass
class RectangleLSP(ShapeLSP):
    width: float
    height: float

    def area(self) -> float:
        return self.width * self.height


@dataclass
class SquareLSP(ShapeLSP):
    side: float

    def area(self) -> float:
        return self.side ** 2


# =============================================================================
# I: Interface Segregation Principle
# Clients should not be forced to depend on interfaces they do not use.
# =============================================================================

# ❌ VIOLATION: SimplePrinter forced to implement scan/fax
class BadMachine(ABC):
    @abstractmethod
    def print_doc(self, doc: str) -> str:
        pass

    @abstractmethod
    def scan(self, doc: str) -> str:
        pass

    @abstractmethod
    def fax(self, doc: str) -> str:
        pass


# ✅ FIX: Segregated interfaces
class Printer(ABC):
    @abstractmethod
    def print_doc(self, doc: str) -> str:
        pass


class Scanner(ABC):
    @abstractmethod
    def scan(self, doc: str) -> str:
        pass


class SimplePrinter(Printer):
    def print_doc(self, doc: str) -> str:
        return f"Printing: {doc}"


class MultiFunctionDevice(Printer, Scanner):
    def print_doc(self, doc: str) -> str:
        return f"Printing: {doc}"

    def scan(self, doc: str) -> str:
        return f"Scanning: {doc}"


# =============================================================================
# D: Dependency Inversion Principle
# High-level modules should depend on abstractions, not concretions.
# =============================================================================

# ❌ VIOLATION: NotificationService directly depends on concrete EmailSender
class EmailSender:
    def send(self, to: str, msg: str) -> str:
        return f"Email to {to}: {msg}"


class BadNotificationService:
    def __init__(self):
        self.sender = EmailSender()  # tight coupling!

    def notify(self, to: str, msg: str) -> str:
        return self.sender.send(to, msg)


# ✅ FIX: Depend on Protocol (abstraction)
class MessageSender(Protocol):
    def send(self, to: str, msg: str) -> str: ...


class EmailMessageSender:
    def send(self, to: str, msg: str) -> str:
        return f"Email to {to}: {msg}"


class SMSMessageSender:
    def send(self, to: str, msg: str) -> str:
        return f"SMS to {to}: {msg}"


class NotificationService:
    def __init__(self, sender: MessageSender):
        self.sender = sender

    def notify(self, to: str, msg: str) -> str:
        return self.sender.send(to, msg)


# =============================================================================
if __name__ == "__main__":
    # S: Single Responsibility
    profile = UserProfile("Alice", "alice@example.com")
    auth = UserAuth("alice@example.com", "hashed_pw")
    profile.update_name("Alice Smith")
    assert profile.name == "Alice Smith"
    assert auth.authenticate("hashed_pw") is True

    # O: Open/Closed
    shapes: list[Shape] = [Circle(5), Rectangle(3, 4)]
    assert abs(area_calculator(shapes) - (math.pi * 25 + 12)) < 0.01

    # L: Liskov Substitution
    def compute_area(shape: ShapeLSP) -> float:
        return shape.area()

    assert compute_area(RectangleLSP(3, 4)) == 12
    assert compute_area(SquareLSP(5)) == 25

    # I: Interface Segregation
    printer = SimplePrinter()
    assert printer.print_doc("report") == "Printing: report"
    mfd = MultiFunctionDevice()
    assert mfd.scan("photo") == "Scanning: photo"

    # D: Dependency Inversion
    email_svc = NotificationService(EmailMessageSender())
    assert "Email" in email_svc.notify("user@test.com", "Hello")
    sms_svc = NotificationService(SMSMessageSender())
    assert "SMS" in sms_svc.notify("+1234567890", "Hello")

    print("All SOLID tests passed!")
