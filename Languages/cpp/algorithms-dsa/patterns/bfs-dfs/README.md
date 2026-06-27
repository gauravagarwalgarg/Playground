# BFS & DFS Graph Traversal Patterns

## When to Identify
- **BFS Keywords**: "shortest path (unweighted)", "minimum steps", "level order", "nearest", "spread/infection"
- **DFS Keywords**: "all paths", "connected components", "cycle detection", "topological sort", "island count"
- **Signals**: Grid problems, tree traversal, graph connectivity, exploring states

## When to Use BFS vs DFS

| Criterion | BFS | DFS |
|-----------|-----|-----|
| Shortest path (unweighted) | ✅ Yes | ❌ No |
| Level-by-level processing | ✅ Yes | ❌ No |
| Nearest/minimum steps | ✅ Yes | ❌ No |
| All paths/solutions | ❌ Expensive | ✅ Yes |
| Cycle detection | ✅ Yes | ✅ Yes |
| Topological sort | ✅ Kahn's algo | ✅ Post-order |
| Memory (wide graph) | ❌ O(width) | ✅ O(depth) |
| Memory (deep graph) | ✅ O(width) | ❌ O(depth)/stack overflow |

## Templates

### BFS (Queue + Visited)
```
queue.push(start), visited[start] = true
while !queue.empty():
    node = queue.front(), queue.pop()
    for neighbor in adj[node]:
        if !visited[neighbor]:
            visited[neighbor] = true
            queue.push(neighbor)
```

### DFS Recursive
```
void dfs(node, visited):
    visited[node] = true
    for neighbor in adj[node]:
        if !visited[neighbor]:
            dfs(neighbor, visited)
```

### DFS Iterative (Stack)
```
stack.push(start)
while !stack.empty():
    node = stack.top(), stack.pop()
    if visited[node]: continue
    visited[node] = true
    for neighbor in adj[node]:
        if !visited[neighbor]:
            stack.push(neighbor)
```

### Multi-Source BFS (Simultaneous Spread)
```
// Push ALL sources into queue initially
for each source: queue.push(source), visited[source] = true
// Standard BFS from here all sources expand simultaneously
```
**Use when**: Rotting Oranges, Walls and Gates, multi-point shortest distance

## Complexity
| Algorithm | Time | Space |
|-----------|------|-------|
| BFS/DFS (adj list) | O(V + E) | O(V) |
| BFS/DFS (grid) | O(rows × cols) | O(rows × cols) |
| Multi-source BFS | O(V + E) | O(V) |

## Key Insight
BFS explores in "rings" outward from source first time you reach a node is the shortest path. DFS explores one branch completely before backtracking natural for exhaustive search and recursion-based problems.
