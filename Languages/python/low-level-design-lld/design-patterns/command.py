"""
Command Pattern

Encapsulates a request as an object, allowing parameterization, queuing,
and undo/redo operations. TextEditor with command history.
"""

from abc import ABC, abstractmethod
from dataclasses import dataclass, field


class Command(ABC):
    @abstractmethod
    def execute(self) -> None:
        pass

    @abstractmethod
    def undo(self) -> None:
        pass


class TextEditor:
    def __init__(self):
        self.content: str = ""


@dataclass
class InsertCommand(Command):
    editor: TextEditor
    text: str
    position: int = 0

    def execute(self) -> None:
        self.editor.content = (
            self.editor.content[: self.position]
            + self.text
            + self.editor.content[self.position:]
        )

    def undo(self) -> None:
        self.editor.content = (
            self.editor.content[: self.position]
            + self.editor.content[self.position + len(self.text):]
        )


@dataclass
class DeleteCommand(Command):
    editor: TextEditor
    position: int
    length: int
    _deleted: str = field(default="", init=False)

    def execute(self) -> None:
        self._deleted = self.editor.content[self.position: self.position + self.length]
        self.editor.content = (
            self.editor.content[: self.position]
            + self.editor.content[self.position + self.length:]
        )

    def undo(self) -> None:
        self.editor.content = (
            self.editor.content[: self.position]
            + self._deleted
            + self.editor.content[self.position:]
        )


class CommandHistory:
    def __init__(self):
        self._history: list[Command] = []
        self._redo_stack: list[Command] = []

    def execute(self, command: Command) -> None:
        command.execute()
        self._history.append(command)
        self._redo_stack.clear()

    def undo(self) -> None:
        if self._history:
            cmd = self._history.pop()
            cmd.undo()
            self._redo_stack.append(cmd)

    def redo(self) -> None:
        if self._redo_stack:
            cmd = self._redo_stack.pop()
            cmd.execute()
            self._history.append(cmd)


if __name__ == "__main__":
    editor = TextEditor()
    history = CommandHistory()

    history.execute(InsertCommand(editor, "Hello", 0))
    assert editor.content == "Hello"

    history.execute(InsertCommand(editor, " World", 5))
    assert editor.content == "Hello World"

    history.undo()
    assert editor.content == "Hello"

    history.redo()
    assert editor.content == "Hello World"

    history.execute(DeleteCommand(editor, 5, 6))
    assert editor.content == "Hello"

    history.undo()
    assert editor.content == "Hello World"

    print("All tests passed!")
