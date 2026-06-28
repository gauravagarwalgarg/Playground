"""
Iterator Pattern

Provides a way to access elements of a collection sequentially without
exposing its underlying representation. Demonstrates Python's iterator protocol.
"""

from dataclasses import dataclass, field
from typing import Iterator


@dataclass
class Task:
    name: str
    priority: int  # lower = higher priority


class TaskQueue:
    """Custom collection that iterates tasks in priority order."""

    def __init__(self):
        self._tasks: list[Task] = []

    def add(self, task: Task) -> None:
        self._tasks.append(task)

    def __len__(self) -> int:
        return len(self._tasks)

    def __iter__(self) -> "TaskIterator":
        """Returns tasks sorted by priority (ascending)."""
        return TaskIterator(sorted(self._tasks, key=lambda t: t.priority))

    def __contains__(self, name: str) -> bool:
        return any(t.name == name for t in self._tasks)


class TaskIterator:
    """Explicit iterator implementing __iter__ and __next__."""

    def __init__(self, tasks: list[Task]):
        self._tasks = tasks
        self._index = 0

    def __iter__(self) -> "TaskIterator":
        return self

    def __next__(self) -> Task:
        if self._index >= len(self._tasks):
            raise StopIteration
        task = self._tasks[self._index]
        self._index += 1
        return task


def filtered_tasks(queue: TaskQueue, max_priority: int) -> Iterator[Task]:
    """Generator-based iterator for filtering."""
    for task in queue:
        if task.priority <= max_priority:
            yield task


if __name__ == "__main__":
    queue = TaskQueue()
    queue.add(Task("Low priority task", 3))
    queue.add(Task("Critical fix", 1))
    queue.add(Task("Medium task", 2))

    # Iterates in priority order
    names = [t.name for t in queue]
    assert names == ["Critical fix", "Medium task", "Low priority task"]

    # Multiple iterations work (new iterator each time)
    assert [t.priority for t in queue] == [1, 2, 3]

    # __contains__ protocol
    assert "Critical fix" in queue
    assert "Nonexistent" not in queue

    # __len__ protocol
    assert len(queue) == 3

    # Generator-based filtering
    urgent = list(filtered_tasks(queue, max_priority=2))
    assert len(urgent) == 2
    assert urgent[0].name == "Critical fix"

    # Manual iteration
    it = iter(queue)
    assert next(it).priority == 1
    assert next(it).priority == 2

    print("All tests passed!")
