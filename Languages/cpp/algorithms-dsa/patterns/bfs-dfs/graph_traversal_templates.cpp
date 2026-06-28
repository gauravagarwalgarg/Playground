/**
 * BFS & DFS Graph Traversal Templates
 * 
 * Core graph traversal patterns for adjacency list and grid problems.
 * Includes: BFS, DFS (recursive + iterative), Multi-Source BFS
 * 
 * Compile: g++ -std=c++17 -o graph_traversal graph_traversal_templates.cpp
 */

#include <iostream>
#include <vector>
#include <queue>
#include <stack>

using namespace std;

// =============================================================================
// TEMPLATE 1: BFS (Breadth-First Search) Shortest Path in Unweighted Graph
// Use when: shortest path, minimum steps, level-order traversal
// Example: Find shortest path from source to all nodes
// =============================================================================
vector<int> bfs(int start, const vector<vector<int>>& adj) {
    int n = adj.size();
    vector<int> dist(n, -1);  // -1 means unvisited
    queue<int> q;
    
    dist[start] = 0;
    q.push(start);
    
    while (!q.empty()) {
        int node = q.front();
        q.pop();
        
        for (int neighbor : adj[node]) {
            if (dist[neighbor] == -1) {  // not visited
                dist[neighbor] = dist[node] + 1;
                q.push(neighbor);
            }
        }
    }
    return dist;  // dist[i] = shortest distance from start to i
}

// =============================================================================
// TEMPLATE 2: BFS Level-by-Level (Track Levels Explicitly)
// Use when: need to process nodes level by level (tree level order, etc.)
// =============================================================================
vector<vector<int>> bfs_levels(int start, const vector<vector<int>>& adj) {
    int n = adj.size();
    vector<bool> visited(n, false);
    vector<vector<int>> levels;
    queue<int> q;
    
    visited[start] = true;
    q.push(start);
    
    while (!q.empty()) {
        int level_size = q.size();  // KEY: process all nodes at current level
        vector<int> current_level;
        
        for (int i = 0; i < level_size; i++) {
            int node = q.front();
            q.pop();
            current_level.push_back(node);
            
            for (int neighbor : adj[node]) {
                if (!visited[neighbor]) {
                    visited[neighbor] = true;
                    q.push(neighbor);
                }
            }
        }
        levels.push_back(current_level);
    }
    return levels;
}

// =============================================================================
// TEMPLATE 3: DFS Recursive
// Use when: exploring all paths, connected components, cycle detection
// Example: Count connected components in undirected graph
// =============================================================================
void dfs_recursive(int node, const vector<vector<int>>& adj, vector<bool>& visited) {
    visited[node] = true;
    
    // Process node here (e.g., add to component)
    
    for (int neighbor : adj[node]) {
        if (!visited[neighbor]) {
            dfs_recursive(neighbor, adj, visited);
        }
    }
}

int countComponents(int n, const vector<vector<int>>& adj) {
    vector<bool> visited(n, false);
    int components = 0;
    
    for (int i = 0; i < n; i++) {
        if (!visited[i]) {
            dfs_recursive(i, adj, visited);
            components++;
        }
    }
    return components;
}

// =============================================================================
// TEMPLATE 4: DFS Iterative (Stack-based)
// Use when: avoiding recursion stack overflow for deep graphs
// Note: visits nodes in different order than recursive DFS
// =============================================================================
void dfs_iterative(int start, const vector<vector<int>>& adj, vector<bool>& visited) {
    stack<int> stk;
    stk.push(start);
    
    while (!stk.empty()) {
        int node = stk.top();
        stk.pop();
        
        if (visited[node]) continue;  // skip if already processed
        visited[node] = true;
        
        // Process node here
        
        for (int neighbor : adj[node]) {
            if (!visited[neighbor]) {
                stk.push(neighbor);
            }
        }
    }
}

// =============================================================================
// TEMPLATE 5: Multi-Source BFS
// Use when: spreading from multiple sources simultaneously
// Example: Rotting Oranges find time for all oranges to rot
// Grid version with directions
// =============================================================================
int multiSourceBFS(vector<vector<int>>& grid) {
    int rows = grid.size(), cols = grid[0].size();
    queue<pair<int,int>> q;
    
    // Step 1: Push ALL sources into the queue
    for (int r = 0; r < rows; r++) {
        for (int c = 0; c < cols; c++) {
            if (grid[r][c] == 2) {  // 2 = rotten orange (source)
                q.push({r, c});
            }
        }
    }
    
    // Step 2: Standard BFS from all sources simultaneously
    int directions[4][2] = {{0,1}, {0,-1}, {1,0}, {-1,0}};
    int time = 0;
    
    while (!q.empty()) {
        int level_size = q.size();
        bool rotted_any = false;
        
        for (int i = 0; i < level_size; i++) {
            auto [r, c] = q.front();
            q.pop();
            
            for (auto& dir : directions) {
                int nr = r + dir[0], nc = c + dir[1];
                
                if (nr >= 0 && nr < rows && nc >= 0 && nc < cols 
                    && grid[nr][nc] == 1) {  // 1 = fresh orange
                    grid[nr][nc] = 2;  // mark as rotten (visited)
                    q.push({nr, nc});
                    rotted_any = true;
                }
            }
        }
        if (rotted_any) time++;
    }
    return time;
}

// =============================================================================
// TEMPLATE 6: Grid BFS (Common Pattern)
// Use when: finding shortest path in a grid with obstacles
// =============================================================================
int shortestPathGrid(vector<vector<int>>& grid, pair<int,int> start, pair<int,int> end) {
    int rows = grid.size(), cols = grid[0].size();
    int directions[4][2] = {{0,1}, {0,-1}, {1,0}, {-1,0}};
    
    vector<vector<bool>> visited(rows, vector<bool>(cols, false));
    queue<pair<int,int>> q;
    
    q.push(start);
    visited[start.first][start.second] = true;
    int steps = 0;
    
    while (!q.empty()) {
        int level_size = q.size();
        
        for (int i = 0; i < level_size; i++) {
            auto [r, c] = q.front();
            q.pop();
            
            if (r == end.first && c == end.second) return steps;
            
            for (auto& dir : directions) {
                int nr = r + dir[0], nc = c + dir[1];
                if (nr >= 0 && nr < rows && nc >= 0 && nc < cols
                    && !visited[nr][nc] && grid[nr][nc] == 0) {
                    visited[nr][nc] = true;
                    q.push({nr, nc});
                }
            }
        }
        steps++;
    }
    return -1;  // no path found
}

// =============================================================================
int main() {
    // Build a sample graph: 0-1-2-3, 0-4, 1-4
    int n = 5;
    vector<vector<int>> adj(n);
    adj[0] = {1, 4};
    adj[1] = {0, 2, 4};
    adj[2] = {1, 3};
    adj[3] = {2};
    adj[4] = {0, 1};
    
    // BFS shortest distance from node 0
    vector<int> dist = bfs(0, adj);
    cout << "BFS distances from 0: ";
    for (int d : dist) cout << d << " ";
    cout << endl;
    
    // BFS levels
    auto levels = bfs_levels(0, adj);
    cout << "BFS levels:" << endl;
    for (int i = 0; i < (int)levels.size(); i++) {
        cout << "  Level " << i << ": ";
        for (int node : levels[i]) cout << node << " ";
        cout << endl;
    }
    
    // Connected components
    cout << "Connected components: " << countComponents(n, adj) << endl;
    
    // Multi-source BFS (rotting oranges)
    vector<vector<int>> grid = {{2,1,1},{1,1,0},{0,1,1}};
    cout << "Time to rot all oranges: " << multiSourceBFS(grid) << endl;
    
    return 0;
}
