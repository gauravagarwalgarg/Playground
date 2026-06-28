"""
Low-Level Design: Task Scheduler

Priority queue-based scheduler. Tasks have priority and deadline.
Lower priority number = higher urgency. Ties broken by earliest deadline.
"""

from dataclasses import dataclass, field
from datetime import datetime
import heapq


@dataclass(order=True)
class Task:
    priority: int
    deadline: datetime = field(compare=True)
    name: str = field(compare=False)
    created_at: datetime = field(default_factory=datetime.now, compare=False)


class TaskScheduler:
    def __init__(self):
        self._queue: list[Task] = []
        self._completed: list[Task] = []

    def add_task(self, name: str, priority: int, deadline: datetime) -> Task:
        task = Task(priority=priority, deadline=deadline, name=name)
        heapq.heappush(self._queue, task)
        return task

    def execute_next(self) -> Task | None:
        if not self._queue:
            return None
        task = heapq.heappop(self._queue)
        self._completed.append(task)
        return task

    def peek(self) -> Task | None:
        return self._queue[0] if self._queue else None

    @property
    def pending_count(self) -> int:
        return len(self._queue)

    @property
    def completed_count(self) -> int:
        return len(self._completed)


if __name__ == "__main__":
    scheduler = TaskScheduler()

    d1 = datetime(2024, 12, 1)
    d2 = datetime(2024, 12, 5)
    d3 = datetime(2024, 12, 3)

    scheduler.add_task("Low priority", priority=3, deadline=d1)
    scheduler.add_task("Urgent fix", priority=1, deadline=d2)
    scheduler.add_task("Medium task", priority=2, deadline=d3)

    assert scheduler.pending_count == 3

    # Highest priority first
    task = scheduler.execute_next()
    assert task.name == "Urgent fix"
    assert task.priority == 1

    task = scheduler.execute_next()
    assert task.name == "Medium task"

    task = scheduler.execute_next()
    assert task.name == "Low priority"

    assert scheduler.pending_count == 0
    assert scheduler.completed_count == 3
    assert scheduler.execute_next() is None

    # Same priority: earlier deadline first
    scheduler2 = TaskScheduler()
    scheduler2.add_task("Later", priority=1, deadline=datetime(2024, 12, 10))
    scheduler2.add_task("Sooner", priority=1, deadline=datetime(2024, 12, 1))
    assert scheduler2.execute_next().name == "Sooner"

    print("All tests passed!")
