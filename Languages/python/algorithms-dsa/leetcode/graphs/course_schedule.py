"""
LeetCode #207 - Course Schedule
Topic: Graphs
Difficulty: Medium

Determine if all courses can be finished using topological sort (cycle detection).

Time Complexity: O(V + E)
Space Complexity: O(V + E)
"""
from collections import deque


def can_finish(num_courses: int, prerequisites: list[list[int]]) -> bool:
    graph: list[list[int]] = [[] for _ in range(num_courses)]
    in_degree = [0] * num_courses
    for course, prereq in prerequisites:
        graph[prereq].append(course)
        in_degree[course] += 1
    queue = deque(i for i in range(num_courses) if in_degree[i] == 0)
    visited = 0
    while queue:
        node = queue.popleft()
        visited += 1
        for neighbor in graph[node]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)
    return visited == num_courses


if __name__ == "__main__":
    assert can_finish(2, [[1, 0]]) is True
    assert can_finish(2, [[1, 0], [0, 1]]) is False
    assert can_finish(3, [[1, 0], [2, 1]]) is True
    assert can_finish(1, []) is True
    print("All tests passed!")
