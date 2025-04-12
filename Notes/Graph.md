# Graph

- Edge and nodes

- Every degree has 2 * Edges

## Store

- `n` nodes, `m` edges
- m lines represent edges

[[2 1], [1, 3], [2, 4], [3, 4], [2, 5], [4, 5]]

- Matrix

adj[n + 1][n + 1]

adjacentCMatrix


```cpp

int n, m;
int adj[n + 1][m + 1]

for(int i = 0; i < m; i++) {
		int u, v;
		cin >> u >> v;
		adj[u][v] = 1;
		adj[v][u] = 1;
}
```

- list
adjacent neighbors

vector<int> adj[n + 1]

```cpp
int n, m;
vector<int> adj[n + 1]

for(int i = 0; i < m; i++) {
		int u, v;
		cin >> u >> v;
		adj[u].push_back(v);
}

```

- Weighted Graph

```
For matrix, store `weight` instead of `1`
For list, store it in pair {neighbour, weight}
```

# Connected Components


