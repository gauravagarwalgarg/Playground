"""
Factory Method Pattern

Defines an interface for creating objects, letting subclasses decide which
class to instantiate. New shapes can be added without modifying ShapeFactory
by registering them at runtime.
"""

from abc import ABC, abstractmethod
import math


class Shape(ABC):
    @abstractmethod
    def area(self) -> float:
        pass

    @abstractmethod
    def describe(self) -> str:
        pass


class Circle(Shape):
    def __init__(self, radius: float):
        self.radius = radius

    def area(self) -> float:
        return math.pi * self.radius ** 2

    def describe(self) -> str:
        return f"Circle(radius={self.radius})"


class Rectangle(Shape):
    def __init__(self, width: float, height: float):
        self.width = width
        self.height = height

    def area(self) -> float:
        return self.width * self.height

    def describe(self) -> str:
        return f"Rectangle({self.width}x{self.height})"


class Triangle(Shape):
    def __init__(self, base: float, height: float):
        self.base = base
        self.height = height

    def area(self) -> float:
        return 0.5 * self.base * self.height

    def describe(self) -> str:
        return f"Triangle(base={self.base}, height={self.height})"


class ShapeFactory:
    """Open for extension: register new shapes without modifying factory code."""

    _registry: dict[str, type[Shape]] = {}

    @classmethod
    def register(cls, name: str, shape_cls: type[Shape]) -> None:
        cls._registry[name] = shape_cls

    @classmethod
    def create(cls, name: str, **kwargs) -> Shape:
        if name not in cls._registry:
            raise ValueError(f"Unknown shape: {name}")
        return cls._registry[name](**kwargs)


# Register built-in shapes
ShapeFactory.register("circle", Circle)
ShapeFactory.register("rectangle", Rectangle)
ShapeFactory.register("triangle", Triangle)


if __name__ == "__main__":
    circle = ShapeFactory.create("circle", radius=5)
    assert abs(circle.area() - 78.54) < 0.01
    assert circle.describe() == "Circle(radius=5)"

    rect = ShapeFactory.create("rectangle", width=3, height=4)
    assert rect.area() == 12.0

    tri = ShapeFactory.create("triangle", base=6, height=4)
    assert tri.area() == 12.0

    # Extend without modifying factory
    class Pentagon(Shape):
        def __init__(self, side: float):
            self.side = side

        def area(self) -> float:
            return (math.sqrt(5 * (5 + 2 * math.sqrt(5))) / 4) * self.side ** 2

        def describe(self) -> str:
            return f"Pentagon(side={self.side})"

    ShapeFactory.register("pentagon", Pentagon)
    pent = ShapeFactory.create("pentagon", side=3)
    assert pent.area() > 0

    print("All tests passed!")
