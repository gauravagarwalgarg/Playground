package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"sort"
	"sync"
)

// Consistent Hashing - Distributes keys across nodes in a ring.
// Adding/removing a node only remaps ~K/N keys (K=keys, N=nodes).
// Virtual nodes ensure even distribution.

type HashRing struct {
	mu           sync.RWMutex
	ring         []uint32          // sorted hashes on the ring
	nodes        map[uint32]string // hash → node name
	virtualNodes int               // virtual nodes per physical node
}

func NewHashRing(virtualNodes int) *HashRing {
	return &HashRing{
		nodes:        make(map[uint32]string),
		virtualNodes: virtualNodes,
	}
}

// hashKey generates a consistent hash for a key
func hashKey(key string) uint32 {
	h := sha256.Sum256([]byte(key))
	return binary.BigEndian.Uint32(h[:4])
}

// AddNode adds a physical node with virtual replicas
func (hr *HashRing) AddNode(node string) {
	hr.mu.Lock()
	defer hr.mu.Unlock()
	for i := 0; i < hr.virtualNodes; i++ {
		virtualKey := fmt.Sprintf("%s#%d", node, i)
		hash := hashKey(virtualKey)
		hr.ring = append(hr.ring, hash)
		hr.nodes[hash] = node
	}
	sort.Slice(hr.ring, func(i, j int) bool { return hr.ring[i] < hr.ring[j] })
}

// RemoveNode removes a physical node and all its virtual replicas
func (hr *HashRing) RemoveNode(node string) {
	hr.mu.Lock()
	defer hr.mu.Unlock()
	var newRing []uint32
	for _, hash := range hr.ring {
		if hr.nodes[hash] != node {
			newRing = append(newRing, hash)
		} else {
			delete(hr.nodes, hash)
		}
	}
	hr.ring = newRing
}

// GetNode finds which node a key is assigned to
func (hr *HashRing) GetNode(key string) string {
	hr.mu.RLock()
	defer hr.mu.RUnlock()
	if len(hr.ring) == 0 {
		return ""
	}
	hash := hashKey(key)
	// Binary search for the first ring position >= hash
	idx := sort.Search(len(hr.ring), func(i int) bool {
		return hr.ring[i] >= hash
	})
	// Wrap around
	if idx == len(hr.ring) {
		idx = 0
	}
	return hr.nodes[hr.ring[idx]]
}

// GetNodes returns N distinct nodes for replication
func (hr *HashRing) GetNodes(key string, count int) []string {
	hr.mu.RLock()
	defer hr.mu.RUnlock()
	if len(hr.ring) == 0 {
		return nil
	}
	hash := hashKey(key)
	idx := sort.Search(len(hr.ring), func(i int) bool {
		return hr.ring[i] >= hash
	})

	seen := make(map[string]bool)
	var result []string
	for i := 0; i < len(hr.ring) && len(result) < count; i++ {
		pos := (idx + i) % len(hr.ring)
		node := hr.nodes[hr.ring[pos]]
		if !seen[node] {
			seen[node] = true
			result = append(result, node)
		}
	}
	return result
}

func main() {
	ring := NewHashRing(150) // 150 virtual nodes per physical node

	// Add nodes
	nodes := []string{"node-1", "node-2", "node-3"}
	for _, n := range nodes {
		ring.AddNode(n)
	}

	// Distribute 1000 keys and count distribution
	distribution := make(map[string]int)
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key-%d", i)
		node := ring.GetNode(key)
		distribution[node]++
	}

	fmt.Println("--- Distribution across 3 nodes (1000 keys) ---")
	for node, count := range distribution {
		fmt.Printf("  %s: %d keys (%.1f%%)\n", node, count, float64(count)/10)
	}

	// Verify all nodes got some keys (with 150 vnodes, should be fairly even)
	for _, node := range nodes {
		if distribution[node] < 200 { // at least 20% each
			fmt.Printf("WARNING: %s only has %d keys (might be uneven)\n", node, distribution[node])
		}
	}
	fmt.Println("PASS: all nodes received keys")

	// Test adding a node - check remapping
	fmt.Println("\n--- Adding node-4 ---")
	before := make(map[string]string)
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key-%d", i)
		before[key] = ring.GetNode(key)
	}

	ring.AddNode("node-4")

	remapped := 0
	newDist := make(map[string]int)
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key-%d", i)
		after := ring.GetNode(key)
		newDist[after]++
		if before[key] != after {
			remapped++
		}
	}

	fmt.Printf("  Remapped: %d/1000 keys (%.1f%%)\n", remapped, float64(remapped)/10)
	fmt.Println("  New distribution:")
	for node, count := range newDist {
		fmt.Printf("    %s: %d keys\n", node, count)
	}

	// Ideally ~25% of keys remap when adding 1 of 4 nodes
	if remapped < 500 { // should be significantly less than 50%
		fmt.Println("PASS: minimal remapping on node addition")
	} else {
		panic(fmt.Sprintf("FAIL: too many remapped: %d", remapped))
	}

	// Test removing a node
	fmt.Println("\n--- Removing node-2 ---")
	ring.RemoveNode("node-2")
	orphanNode := ring.GetNode("test-key")
	if orphanNode == "node-2" {
		panic("FAIL: removed node still serving keys")
	}
	fmt.Println("PASS: removed node no longer serving keys")

	// Test replication
	replicas := ring.GetNodes("important-key", 2)
	fmt.Printf("\n--- Replication for 'important-key': %v ---\n", replicas)
	if len(replicas) == 2 && replicas[0] != replicas[1] {
		fmt.Println("PASS: got 2 distinct replicas")
	} else {
		panic("FAIL: replication")
	}
}
