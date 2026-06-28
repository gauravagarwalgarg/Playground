"""
Strategy Pattern

Defines a family of algorithms, encapsulates each one, and makes them
interchangeable. PaymentProcessor can swap strategies at runtime.
"""

from abc import ABC, abstractmethod
from dataclasses import dataclass


class PaymentStrategy(ABC):
    @abstractmethod
    def pay(self, amount: float) -> str:
        pass


class CreditCardPayment(PaymentStrategy):
    def __init__(self, card_number: str):
        self.card_number = card_number

    def pay(self, amount: float) -> str:
        return f"Paid ${amount:.2f} via Credit Card ending {self.card_number[-4:]}"


class PayPalPayment(PaymentStrategy):
    def __init__(self, email: str):
        self.email = email

    def pay(self, amount: float) -> str:
        return f"Paid ${amount:.2f} via PayPal ({self.email})"


class CryptoPayment(PaymentStrategy):
    def __init__(self, wallet: str):
        self.wallet = wallet

    def pay(self, amount: float) -> str:
        return f"Paid ${amount:.2f} via Crypto wallet {self.wallet[:8]}..."


@dataclass
class PaymentProcessor:
    strategy: PaymentStrategy

    def process(self, amount: float) -> str:
        return self.strategy.pay(amount)

    def set_strategy(self, strategy: PaymentStrategy) -> None:
        self.strategy = strategy


if __name__ == "__main__":
    processor = PaymentProcessor(CreditCardPayment("4111222233334444"))
    result = processor.process(99.99)
    assert "Credit Card" in result
    assert "4444" in result

    # Swap strategy at runtime
    processor.set_strategy(PayPalPayment("user@example.com"))
    result = processor.process(50.00)
    assert "PayPal" in result
    assert "user@example.com" in result

    processor.set_strategy(CryptoPayment("0xABCDEF1234567890"))
    result = processor.process(200.00)
    assert "Crypto" in result
    assert "0xABCDEF" in result

    print("All tests passed!")
