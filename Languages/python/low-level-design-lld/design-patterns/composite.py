"""
Composite Pattern

Composes objects into tree structures. File and Directory share a common
interface, letting clients treat individual objects and compositions uniformly.
"""

from abc import ABC, abstractmethod


class FileSystemNode(ABC):
    def __init__(self, name: str):
        self.name = name

    @abstractmethod
    def size(self) -> int:
        pass

    @abstractmethod
    def display(self, indent: int = 0) -> str:
        pass


class File(FileSystemNode):
    def __init__(self, name: str, size: int):
        super().__init__(name)
        self._size = size

    def size(self) -> int:
        return self._size

    def display(self, indent: int = 0) -> str:
        return f"{'  ' * indent}{self.name} ({self._size}B)"


class Directory(FileSystemNode):
    def __init__(self, name: str):
        super().__init__(name)
        self.children: list[FileSystemNode] = []

    def add(self, node: FileSystemNode) -> "Directory":
        self.children.append(node)
        return self

    def remove(self, name: str) -> None:
        self.children = [c for c in self.children if c.name != name]

    def size(self) -> int:
        return sum(child.size() for child in self.children)

    def display(self, indent: int = 0) -> str:
        lines = [f"{'  ' * indent}{self.name}/"]
        for child in self.children:
            lines.append(child.display(indent + 1))
        return "\n".join(lines)

    def find(self, name: str) -> FileSystemNode | None:
        for child in self.children:
            if child.name == name:
                return child
            if isinstance(child, Directory):
                found = child.find(name)
                if found:
                    return found
        return None


if __name__ == "__main__":
    root = Directory("root")
    src = Directory("src")
    src.add(File("main.py", 1200))
    src.add(File("utils.py", 800))
    root.add(src)
    root.add(File("README.md", 500))

    # Uniform interface: size works on both
    assert root.size() == 2500
    assert src.size() == 2000

    # Display tree
    output = root.display()
    assert "root/" in output
    assert "  src/" in output
    assert "    main.py (1200B)" in output

    # Find across tree
    assert root.find("utils.py") is not None
    assert root.find("nonexistent") is None

    # Remove
    src.remove("utils.py")
    assert src.size() == 1200
    assert root.size() == 1700

    print("All tests passed!")
