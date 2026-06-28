import java.util.*;

/**
 * Graph Patterns
 * 
 * Key patterns:
 * 1. BFS Shortest Path: Level-order traversal for unweighted graphs
 * 2. Topological Sort (Kahn's): Process nodes with 0 in-degree first
 * 3. Union-Find (Disjoint Set): Path compression + union by rank
 * 4. Dijkstra's Algorithm: Greedy shortest path for weighted graphs
 * 
 * When to use which:
 * - BFS: Unweighted shortest path, level-order processing
 * - Topological Sort: DAG ordering, dependency resolution, cycle detection
 * - Union-Find: Dynamic connectivity, cycle detection in undirected graphs
 * - Dijkstra: Weighted shortest path (non-negative weights)
 */
public class Graphs {

    // ==================== BFS Shortest Path ====================
    /**
     * Find shortest path in an unweighted graph using BFS.
     * Each edge has weight 1, so the first time BFS visits a node is the shortest path.
     * 
     * Time: O(V + E), Space: O(V)
     */
    public static int bfsShortestPath(Map<Integer, List<Integer>> graph, int source, int target) {
        if (source == target) return 0;

        Queue<Integer> queue = new LinkedList<>();
        Set<Integer> visited = new HashSet<>();

        queue.offer(source);
        visited.add(source);
        int distance = 0;

        while (!queue.isEmpty()) {
            distance++;
            int size = queue.size();

            // Process entire level
            for (int i = 0; i < size; i++) {
                int node = queue.poll();
                for (int neighbor : graph.getOrDefault(node, Collections.emptyList())) {
                    if (neighbor == target) return distance;
                    if (!visited.contains(neighbor)) {
                        visited.add(neighbor);
                        queue.offer(neighbor);
                    }
                }
            }
        }

        return -1; // Target not reachable
    }

    // ==================== Topological Sort (Kahn's Algorithm) ====================
    /**
     * Kahn's Algorithm: BFS-based topological sort.
     * Strategy: Start with all nodes that have 0 in-degree.
     * Process them, reduce neighbors' in-degree, add new 0-degree nodes.
     * If result size < numNodes, there's a cycle.
     * 
     * Time: O(V + E), Space: O(V + E)
     */
    public static List<Integer> topologicalSort(int numNodes, int[][] edges) {
        // Build adjacency list and in-degree count
        Map<Integer, List<Integer>> graph = new HashMap<>();
        int[] inDegree = new int[numNodes];

        for (int[] edge : edges) {
            graph.computeIfAbsent(edge[0], k -> new ArrayList<>()).add(edge[1]);
            inDegree[edge[1]]++;
        }

        // Start with all nodes having 0 in-degree
        Queue<Integer> queue = new LinkedList<>();
        for (int i = 0; i < numNodes; i++) {
            if (inDegree[i] == 0) {
                queue.offer(i);
            }
        }

        List<Integer> result = new ArrayList<>();
        while (!queue.isEmpty()) {
            int node = queue.poll();
            result.add(node);

            for (int neighbor : graph.getOrDefault(node, Collections.emptyList())) {
                inDegree[neighbor]--;
                if (inDegree[neighbor] == 0) {
                    queue.offer(neighbor);
                }
            }
        }

        // If not all nodes are processed, graph has a cycle
        return result.size() == numNodes ? result : Collections.emptyList();
    }

    // ==================== Union-Find with Path Compression + Rank ====================
    /**
     * Union-Find (Disjoint Set Union) data structure.
     * - find(): Path compression flattens the tree for O(α(n)) amortized
     * - union(): Union by rank keeps tree balanced
     * 
     * Applications: Connected components, cycle detection, Kruskal's MST
     */
    static class UnionFind {
        private int[] parent;
        private int[] rank;
        private int components;

        public UnionFind(int n) {
            parent = new int[n];
            rank = new int[n];
            components = n;
            for (int i = 0; i < n; i++) {
                parent[i] = i; // Each node is its own parent initially
            }
        }

        /**
         * Find with path compression.
         * Makes every node on the path point directly to the root.
         */
        public int find(int x) {
            if (parent[x] != x) {
                parent[x] = find(parent[x]); // Path compression
            }
            return parent[x];
        }

        /**
         * Union by rank.
         * Attach the shorter tree under the taller tree's root.
         * Returns true if a merge happened (they were in different sets).
         */
        public boolean union(int x, int y) {
            int rootX = find(x);
            int rootY = find(y);

            if (rootX == rootY) return false; // Already connected

            // Union by rank
            if (rank[rootX] < rank[rootY]) {
                parent[rootX] = rootY;
            } else if (rank[rootX] > rank[rootY]) {
                parent[rootY] = rootX;
            } else {
                parent[rootY] = rootX;
                rank[rootX]++;
            }

            components--;
            return true;
        }

        public boolean connected(int x, int y) {
            return find(x) == find(y);
        }

        public int getComponents() {
            return components;
        }
    }

    // ==================== Dijkstra's Algorithm ====================
    /**
     * Dijkstra's shortest path for weighted graphs (non-negative weights).
     * Strategy: Greedy BFS using a min-heap ordered by distance.
     * Always process the closest unvisited node first.
     * 
     * Time: O((V + E) log V), Space: O(V + E)
     * 
     * @param graph adjacency list: node -> list of [neighbor, weight]
     * @param source starting node
     * @param n number of nodes
     * @return shortest distance from source to every node (-1 if unreachable)
     */
    public static int[] dijkstra(Map<Integer, List<int[]>> graph, int source, int n) {
        int[] dist = new int[n];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[source] = 0;

        // Min-heap: [distance, node]
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]);
        pq.offer(new int[]{0, source});

        while (!pq.isEmpty()) {
            int[] curr = pq.poll();
            int d = curr[0], node = curr[1];

            // Skip if we already found a shorter path
            if (d > dist[node]) continue;

            for (int[] edge : graph.getOrDefault(node, Collections.emptyList())) {
                int neighbor = edge[0], weight = edge[1];
                int newDist = dist[node] + weight;

                if (newDist < dist[neighbor]) {
                    dist[neighbor] = newDist;
                    pq.offer(new int[]{newDist, neighbor});
                }
            }
        }

        // Convert MAX_VALUE to -1 for unreachable nodes
        for (int i = 0; i < n; i++) {
            if (dist[i] == Integer.MAX_VALUE) dist[i] = -1;
        }

        return dist;
    }

    // ==================== Tests ====================
    public static void main(String[] args) {
        // Test BFS Shortest Path
        // Graph: 0 -- 1 -- 3
        //        |         |
        //        2 ------- 4
        Map<Integer, List<Integer>> bfsGraph = new HashMap<>();
        bfsGraph.put(0, Arrays.asList(1, 2));
        bfsGraph.put(1, Arrays.asList(0, 3));
        bfsGraph.put(2, Arrays.asList(0, 4));
        bfsGraph.put(3, Arrays.asList(1, 4));
        bfsGraph.put(4, Arrays.asList(2, 3));

        assert bfsShortestPath(bfsGraph, 0, 4) == 2 : "BFS test 1 failed";
        assert bfsShortestPath(bfsGraph, 0, 3) == 2 : "BFS test 2 failed";
        assert bfsShortestPath(bfsGraph, 0, 0) == 0 : "BFS test 3 failed";
        assert bfsShortestPath(bfsGraph, 1, 4) == 2 : "BFS test 4 failed";

        // Test Topological Sort
        // DAG: 0 -> 1, 0 -> 2, 1 -> 3, 2 -> 3
        int[][] edges = {{0, 1}, {0, 2}, {1, 3}, {2, 3}};
        List<Integer> topoResult = topologicalSort(4, edges);
        assert topoResult.size() == 4 : "Topo sort test 1 failed: not all nodes";
        // Verify ordering: for each edge u->v, u appears before v
        Map<Integer, Integer> position = new HashMap<>();
        for (int i = 0; i < topoResult.size(); i++) {
            position.put(topoResult.get(i), i);
        }
        for (int[] edge : edges) {
            assert position.get(edge[0]) < position.get(edge[1])
                : "Topo sort order violated for edge " + edge[0] + " -> " + edge[1];
        }

        // Test cycle detection (topological sort returns empty)
        int[][] cycleEdges = {{0, 1}, {1, 2}, {2, 0}};
        List<Integer> cycleResult = topologicalSort(3, cycleEdges);
        assert cycleResult.isEmpty() : "Topo sort cycle detection failed";

        // Test Union-Find
        UnionFind uf = new UnionFind(5);
        assert uf.getComponents() == 5 : "UF init failed";
        uf.union(0, 1);
        uf.union(2, 3);
        assert uf.getComponents() == 3 : "UF components test 1 failed";
        assert uf.connected(0, 1) : "UF connected test 1 failed";
        assert !uf.connected(0, 2) : "UF connected test 2 failed";
        uf.union(1, 3);
        assert uf.connected(0, 3) : "UF connected test 3 failed (transitive)";
        assert uf.getComponents() == 2 : "UF components test 2 failed";

        // Test Dijkstra
        // Weighted graph: 0 --(1)--> 1 --(2)--> 3
        //                 0 --(4)--> 2 --(1)--> 3
        Map<Integer, List<int[]>> weightedGraph = new HashMap<>();
        weightedGraph.put(0, Arrays.asList(new int[]{1, 1}, new int[]{2, 4}));
        weightedGraph.put(1, Arrays.asList(new int[]{3, 2}));
        weightedGraph.put(2, Arrays.asList(new int[]{3, 1}));
        weightedGraph.put(3, new ArrayList<>());

        int[] distances = dijkstra(weightedGraph, 0, 4);
        assert distances[0] == 0 : "Dijkstra test: dist to self failed";
        assert distances[1] == 1 : "Dijkstra test: dist to 1 failed";
        assert distances[2] == 4 : "Dijkstra test: dist to 2 failed";
        assert distances[3] == 3 : "Dijkstra test: dist to 3 failed (via 1)";

        // Dijkstra with unreachable node
        Map<Integer, List<int[]>> disconnected = new HashMap<>();
        disconnected.put(0, Arrays.asList(new int[]{1, 5}));
        disconnected.put(1, new ArrayList<>());
        int[] dist2 = dijkstra(disconnected, 0, 3);
        assert dist2[2] == -1 : "Dijkstra unreachable test failed";

        System.out.println("All graph pattern tests passed!");
    }
}
