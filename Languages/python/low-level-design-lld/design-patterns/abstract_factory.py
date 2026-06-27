"""
Abstract Factory Pattern

Provides an interface for creating families of related objects without
specifying their concrete classes. Each factory produces a consistent
theme (Mac/Windows) of UI components.
"""

from abc import ABC, abstractmethod


class Button(ABC):
    @abstractmethod
    def render(self) -> str:
        pass


class Checkbox(ABC):
    @abstractmethod
    def render(self) -> str:
        pass


class MacButton(Button):
    def render(self) -> str:
        return "[ Mac Button ]"


class MacCheckbox(Checkbox):
    def render(self) -> str:
        return "[✓] Mac Checkbox"


class WindowsButton(Button):
    def render(self) -> str:
        return "[ Win Button ]"


class WindowsCheckbox(Checkbox):
    def render(self) -> str:
        return "[☑] Win Checkbox"


class UIFactory(ABC):
    @abstractmethod
    def create_button(self) -> Button:
        pass

    @abstractmethod
    def create_checkbox(self) -> Checkbox:
        pass


class MacFactory(UIFactory):
    def create_button(self) -> Button:
        return MacButton()

    def create_checkbox(self) -> Checkbox:
        return MacCheckbox()


class WindowsFactory(UIFactory):
    def create_button(self) -> Button:
        return WindowsButton()

    def create_checkbox(self) -> Checkbox:
        return WindowsCheckbox()


def render_ui(factory: UIFactory) -> list[str]:
    """Client code works with any factory via the abstract interface."""
    button = factory.create_button()
    checkbox = factory.create_checkbox()
    return [button.render(), checkbox.render()]


if __name__ == "__main__":
    mac_ui = render_ui(MacFactory())
    assert mac_ui == ["[ Mac Button ]", "[✓] Mac Checkbox"]

    win_ui = render_ui(WindowsFactory())
    assert win_ui == ["[ Win Button ]", "[☑] Win Checkbox"]

    # Type consistency within a family
    mac = MacFactory()
    assert isinstance(mac.create_button(), MacButton)
    assert isinstance(mac.create_checkbox(), MacCheckbox)

    print("All tests passed!")
