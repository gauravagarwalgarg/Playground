"""
State Pattern

Allows an object to alter its behavior when its internal state changes.
VendingMachine transitions between Idle, HasMoney, and Dispensing states.
"""

from abc import ABC, abstractmethod


class State(ABC):
    @abstractmethod
    def insert_money(self, machine: "VendingMachine") -> str:
        pass

    @abstractmethod
    def select_item(self, machine: "VendingMachine") -> str:
        pass

    @abstractmethod
    def dispense(self, machine: "VendingMachine") -> str:
        pass


class IdleState(State):
    def insert_money(self, machine: "VendingMachine") -> str:
        machine.state = HasMoneyState()
        return "Money accepted"

    def select_item(self, machine: "VendingMachine") -> str:
        return "Insert money first"

    def dispense(self, machine: "VendingMachine") -> str:
        return "Insert money first"


class HasMoneyState(State):
    def insert_money(self, machine: "VendingMachine") -> str:
        return "Money already inserted"

    def select_item(self, machine: "VendingMachine") -> str:
        if machine.inventory > 0:
            machine.state = DispensingState()
            return "Item selected, dispensing..."
        machine.state = IdleState()
        return "Out of stock, returning money"

    def dispense(self, machine: "VendingMachine") -> str:
        return "Select an item first"


class DispensingState(State):
    def insert_money(self, machine: "VendingMachine") -> str:
        return "Please wait, dispensing"

    def select_item(self, machine: "VendingMachine") -> str:
        return "Already dispensing"

    def dispense(self, machine: "VendingMachine") -> str:
        machine.inventory -= 1
        machine.state = IdleState()
        return "Item dispensed!"


class VendingMachine:
    def __init__(self, inventory: int):
        self.inventory = inventory
        self.state: State = IdleState()

    def insert_money(self) -> str:
        return self.state.insert_money(self)

    def select_item(self) -> str:
        return self.state.select_item(self)

    def dispense(self) -> str:
        return self.state.dispense(self)


if __name__ == "__main__":
    vm = VendingMachine(inventory=2)

    assert vm.select_item() == "Insert money first"
    assert vm.insert_money() == "Money accepted"
    assert vm.insert_money() == "Money already inserted"
    assert vm.select_item() == "Item selected, dispensing..."
    assert vm.dispense() == "Item dispensed!"
    assert vm.inventory == 1

    # Full cycle again
    assert vm.insert_money() == "Money accepted"
    assert vm.select_item() == "Item selected, dispensing..."
    assert vm.dispense() == "Item dispensed!"
    assert vm.inventory == 0

    # Out of stock
    assert vm.insert_money() == "Money accepted"
    assert vm.select_item() == "Out of stock, returning money"

    print("All tests passed!")
