package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"sync"
	"sync/atomic"
)

// Load Balancing Algorithms - Common strategies for distributing requests.
// Covers: Round Robin, Weighted Round Robin, Least Connections,
//         Random, IP Hash (Sticky Sessions).

type Server struct {
	Address     string
	Weight      int
	Connections int32
	Healthy     bool
}

// LoadBalancer interface
type LoadBalancer interface {
	Next(clientIP string) *Server
	Name() string
}

// --- Round Robin ---
type RoundRobin struct {
	servers []*Server
	current uint64
}

func NewRoundRobin(servers []*Server) *RoundRobin {
	return &RoundRobin{servers: servers}
}

func (rr *RoundRobin) Name() string { return "RoundRobin" }

func (rr *RoundRobin) Next(_ string) *Server {
	n := uint64(len(rr.servers))
	for attempts := uint64(0); attempts < n; attempts++ {
		idx := atomic.AddUint64(&rr.current, 1) % n
		if rr.servers[idx].Healthy {
			return rr.servers[idx]
		}
	}
	return nil
}

// --- Weighted Round Robin ---
type WeightedRoundRobin struct {
	mu      sync.Mutex
	servers []*Server
	weights []int
	current int
	gcd     int
	maxW    int
	cw      int
}

func NewWeightedRoundRobin(servers []*Server) *WeightedRoundRobin {
	wrr := &WeightedRoundRobin{servers: servers}
	for _, s := range servers {
		wrr.weights = append(wrr.weights, s.Weight)
	}
	wrr.gcd = gcdSlice(wrr.weights)
	wrr.maxW = maxSlice(wrr.weights)
	wrr.current = -1
	return wrr
}

func (wrr *WeightedRoundRobin) Name() string { return "WeightedRoundRobin" }

func (wrr *WeightedRoundRobin) Next(_ string) *Server {
	wrr.mu.Lock()
	defer wrr.mu.Unlock()
	n := len(wrr.servers)
	for {
		wrr.current = (wrr.current + 1) % n
		if wrr.current == 0 {
			wrr.cw -= wrr.gcd
			if wrr.cw <= 0 {
				wrr.cw = wrr.maxW
			}
		}
		s := wrr.servers[wrr.current]
		if s.Healthy && s.Weight >= wrr.cw {
			return s
		}
	}
}

// --- Least Connections ---
type LeastConnections struct {
	servers []*Server
	mu      sync.Mutex
}

func NewLeastConnections(servers []*Server) *LeastConnections {
	return &LeastConnections{servers: servers}
}

func (lc *LeastConnections) Name() string { return "LeastConnections" }

func (lc *LeastConnections) Next(_ string) *Server {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	var best *Server
	for _, s := range lc.servers {
		if !s.Healthy {
			continue
		}
		if best == nil || s.Connections < best.Connections {
			best = s
		}
	}
	if best != nil {
		atomic.AddInt32(&best.Connections, 1)
	}
	return best
}

// --- IP Hash (Sticky Sessions) ---
type IPHash struct {
	servers []*Server
}

func NewIPHash(servers []*Server) *IPHash {
	return &IPHash{servers: servers}
}

func (ih *IPHash) Name() string { return "IPHash" }

func (ih *IPHash) Next(clientIP string) *Server {
	h := fnv.New32a()
	h.Write([]byte(clientIP))
	idx := h.Sum32() % uint32(len(ih.servers))
	if ih.servers[idx].Healthy {
		return ih.servers[idx]
	}
	// Fallback: linear probe
	for i := 1; i < len(ih.servers); i++ {
		next := (int(idx) + i) % len(ih.servers)
		if ih.servers[next].Healthy {
			return ih.servers[next]
		}
	}
	return nil
}

// --- Random ---
type RandomLB struct {
	servers []*Server
}

func NewRandomLB(servers []*Server) *RandomLB {
	return &RandomLB{servers: servers}
}

func (r *RandomLB) Name() string { return "Random" }

func (r *RandomLB) Next(_ string) *Server {
	healthy := make([]*Server, 0)
	for _, s := range r.servers {
		if s.Healthy {
			healthy = append(healthy, s)
		}
	}
	if len(healthy) == 0 {
		return nil
	}
	return healthy[rand.Intn(len(healthy))]
}

// Helpers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdSlice(nums []int) int {
	result := nums[0]
	for _, n := range nums[1:] {
		result = gcd(result, n)
	}
	return result
}

func maxSlice(nums []int) int {
	m := nums[0]
	for _, n := range nums[1:] {
		if n > m {
			m = n
		}
	}
	return m
}

func testLB(lb LoadBalancer, requests int) {
	dist := make(map[string]int)
	for i := 0; i < requests; i++ {
		ip := fmt.Sprintf("192.168.1.%d", i%10)
		server := lb.Next(ip)
		if server != nil {
			dist[server.Address]++
		}
	}
	fmt.Printf("\n[%s] distribution over %d requests:\n", lb.Name(), requests)
	for addr, count := range dist {
		fmt.Printf("  %s: %d (%.1f%%)\n", addr, count, float64(count*100)/float64(requests))
	}
}

func main() {
	servers := []*Server{
		{Address: "10.0.0.1:8080", Weight: 5, Healthy: true},
		{Address: "10.0.0.2:8080", Weight: 3, Healthy: true},
		{Address: "10.0.0.3:8080", Weight: 1, Healthy: true},
	}

	// Test Round Robin
	rr := NewRoundRobin(servers)
	testLB(rr, 90)

	// Test Weighted Round Robin
	wrr := NewWeightedRoundRobin(servers)
	testLB(wrr, 90)

	// Test Least Connections
	// Simulate: server 1 already has load
	servers[0].Connections = 10
	servers[1].Connections = 2
	servers[2].Connections = 0
	lc := NewLeastConnections(servers)
	s := lc.Next("")
	if s.Address != "10.0.0.3:8080" {
		panic("FAIL: LeastConnections should pick server 3")
	}
	fmt.Println("\nPASS: LeastConnections picks least loaded server")

	// Reset connections
	servers[0].Connections = 0
	servers[1].Connections = 0
	servers[2].Connections = 0

	// Test IP Hash - same IP always goes to same server
	ih := NewIPHash(servers)
	s1 := ih.Next("192.168.1.100")
	s2 := ih.Next("192.168.1.100")
	s3 := ih.Next("192.168.1.100")
	if s1.Address != s2.Address || s2.Address != s3.Address {
		panic("FAIL: IPHash should be sticky")
	}
	fmt.Printf("PASS: IPHash sticky for 192.168.1.100 → %s\n", s1.Address)

	// Test health check scenario
	servers[1].Healthy = false
	rr2 := NewRoundRobin(servers)
	for i := 0; i < 10; i++ {
		s := rr2.Next("")
		if s.Address == "10.0.0.2:8080" {
			panic("FAIL: unhealthy server should not be selected")
		}
	}
	fmt.Println("PASS: unhealthy server skipped")

	// Random
	servers[1].Healthy = true
	randomLB := NewRandomLB(servers)
	testLB(randomLB, 90)
	fmt.Println("\nPASS: all load balancing algorithms complete")
}
