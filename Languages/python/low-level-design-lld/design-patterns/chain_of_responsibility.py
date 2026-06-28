"""
Chain of Responsibility Pattern

Passes a request along a chain of handlers. Each handler decides to process
the request or pass it to the next handler. Support ticket escalation.
"""

from abc import ABC, abstractmethod
from dataclasses import dataclass
from enum import IntEnum


class Severity(IntEnum):
    LOW = 1
    MEDIUM = 2
    HIGH = 3
    CRITICAL = 4


@dataclass
class Ticket:
    id: int
    issue: str
    severity: Severity


class SupportHandler(ABC):
    def __init__(self, next_handler: "SupportHandler | None" = None):
        self._next = next_handler

    def handle(self, ticket: Ticket) -> str:
        if self.can_handle(ticket):
            return self.process(ticket)
        if self._next:
            return self._next.handle(ticket)
        return f"Ticket #{ticket.id}: No handler available"

    @abstractmethod
    def can_handle(self, ticket: Ticket) -> bool:
        pass

    @abstractmethod
    def process(self, ticket: Ticket) -> str:
        pass


class Level1Support(SupportHandler):
    def can_handle(self, ticket: Ticket) -> bool:
        return ticket.severity == Severity.LOW

    def process(self, ticket: Ticket) -> str:
        return f"L1 resolved #{ticket.id}: {ticket.issue}"


class Level2Support(SupportHandler):
    def can_handle(self, ticket: Ticket) -> bool:
        return ticket.severity == Severity.MEDIUM

    def process(self, ticket: Ticket) -> str:
        return f"L2 resolved #{ticket.id}: {ticket.issue}"


class ManagerSupport(SupportHandler):
    def can_handle(self, ticket: Ticket) -> bool:
        return ticket.severity in (Severity.HIGH, Severity.CRITICAL)

    def process(self, ticket: Ticket) -> str:
        return f"Manager escalated #{ticket.id}: {ticket.issue}"


def build_support_chain() -> SupportHandler:
    manager = ManagerSupport()
    level2 = Level2Support(next_handler=manager)
    level1 = Level1Support(next_handler=level2)
    return level1


if __name__ == "__main__":
    chain = build_support_chain()

    t1 = Ticket(1, "Password reset", Severity.LOW)
    assert "L1 resolved" in chain.handle(t1)

    t2 = Ticket(2, "App crash on login", Severity.MEDIUM)
    assert "L2 resolved" in chain.handle(t2)

    t3 = Ticket(3, "Data breach detected", Severity.CRITICAL)
    assert "Manager escalated" in chain.handle(t3)

    t4 = Ticket(4, "Server down", Severity.HIGH)
    assert "Manager escalated" in chain.handle(t4)

    print("All tests passed!")
