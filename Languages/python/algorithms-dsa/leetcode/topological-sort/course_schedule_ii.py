"""
LeetCode #210 - Course Schedule II
Topic: Topological Sort
Difficulty: Medium

Return the ordering of courses you should take to finish all courses.
If impossible (cycle), return empty array.

Approach: Kahn's Algorithm (BFS-based topological sort).
1. Build adjacency list and in-degree array.
2. Start BFS from all nodes with in-degree 0.
3. For each processed node, decrement neighbors' in-degree.
4. If result length != numCourses, there's a cycle.

Time Complexity: O(V + E)
Space Complexity: O(V + E)
"""

from collections import deque


def find_order(num_courses: int, prerequisites: list[list[int]]) -> list[int]:
    graph: list[list[int]] = [[] for _ in range(num_courses)]
    in_degree = [0] * num_courses

    for course, prereq in prerequisites:
        graph[prereq].append(course)
        in_degree[course] += 1

    # Start with courses that have no prerequisites
    queue = deque(i for i in range(num_courses) if in_degree[i] == 0)
    order: list[int] = []

    while queue:
        node = queue.popleft()
        order.append(node)
        for neighbor in graph[node]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)

    return order if len(order) == num_courses else []


if __name__ == "__main__":
    assert find_order(2, [[1, 0]]) == [0, 1]
    result = find_order(4, [[1, 0], [2, 0], [3, 1], [3, 2]])
    # Valid topo orders: [0,1,2,3] or [0,2,1,3]
    assert result[0] == 0 and result[-1] == 3
    assert find_order(2, [[1, 0], [0, 1]]) == []  # cycle
    assert find_order(1, []) == [0]
    print("All tests passed!")
