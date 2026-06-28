"""
LeetCode #133 - Clone Graph
Topic: Graphs
Difficulty: Medium

Deep copy a graph using DFS + hashmap to track visited clones.

Time Complexity: O(V + E)
Space Complexity: O(V)
"""


class Node:
    def __init__(self, val: int = 0, neighbors: list["Node"] | None = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []


def clone_graph(node: Node | None) -> Node | None:
    if not node:
        return None
    clones: dict[int, Node] = {}

    def dfs(n: Node) -> Node:
        if n.val in clones:
            return clones[n.val]
        clone = Node(n.val)
        clones[n.val] = clone
        for neighbor in n.neighbors:
            clone.neighbors.append(dfs(neighbor))
        return clone

    return dfs(node)


if __name__ == "__main__":
    # Build graph: 1--2, 1--4, 2--3, 3--4
    n1, n2, n3, n4 = Node(1), Node(2), Node(3), Node(4)
    n1.neighbors = [n2, n4]
    n2.neighbors = [n1, n3]
    n3.neighbors = [n2, n4]
    n4.neighbors = [n1, n3]

    cloned = clone_graph(n1)
    assert cloned is not n1
    assert cloned.val == 1
    assert len(cloned.neighbors) == 2
    assert clone_graph(None) is None
    print("All tests passed!")
