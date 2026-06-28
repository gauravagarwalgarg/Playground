package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"math"
)

// Bloom Filter - Space-efficient probabilistic data structure.
// Can tell you "definitely not in set" or "probably in set".
// False positives possible, false negatives impossible.
// Used in: databases (LSM trees), caches, network routers, spam filters.

type BloomFilter struct {
	bits    []bool
	size    uint
	hashFns int
}

// NewBloomFilter creates a bloom filter optimized for expectedItems with given false positive rate.
func NewBloomFilter(expectedItems int, fpRate float64) *BloomFilter {
	// Optimal size: m = -(n * ln(p)) / (ln2)^2
	m := uint(math.Ceil(-(float64(expectedItems) * math.Log(fpRate)) / (math.Ln2 * math.Ln2)))
	// Optimal hash functions: k = (m/n) * ln2
	k := int(math.Ceil((float64(m) / float64(expectedItems)) * math.Ln2))

	return &BloomFilter{
		bits:    make([]bool, m),
		size:    m,
		hashFns: k,
	}
}

// hash generates k different hash values for a key
func (bf *BloomFilter) hash(key string) []uint {
	positions := make([]uint, bf.hashFns)

	// Use two hash functions to simulate k hash functions
	// h(i) = h1 + i*h2
	h := fnv.New64a()
	h.Write([]byte(key))
	h1 := h.Sum64()

	s := sha256.Sum256([]byte(key))
	h2 := binary.BigEndian.Uint64(s[:8])

	for i := 0; i < bf.hashFns; i++ {
		positions[i] = uint((h1 + uint64(i)*h2) % uint64(bf.size))
	}
	return positions
}

// Add inserts a key into the bloom filter
func (bf *BloomFilter) Add(key string) {
	for _, pos := range bf.hash(key) {
		bf.bits[pos] = true
	}
}

// Contains checks if a key might be in the set
// Returns false = definitely not in set
// Returns true = probably in set (may be false positive)
func (bf *BloomFilter) Contains(key string) bool {
	for _, pos := range bf.hash(key) {
		if !bf.bits[pos] {
			return false
		}
	}
	return true
}

// FillRatio returns the proportion of bits set
func (bf *BloomFilter) FillRatio() float64 {
	set := 0
	for _, b := range bf.bits {
		if b {
			set++
		}
	}
	return float64(set) / float64(bf.size)
}

func (bf *BloomFilter) Stats() {
	fmt.Printf("  Size: %d bits (%.2f KB)\n", bf.size, float64(bf.size)/8/1024)
	fmt.Printf("  Hash functions: %d\n", bf.hashFns)
	fmt.Printf("  Fill ratio: %.4f\n", bf.FillRatio())
}

func main() {
	// Create a bloom filter for 10000 items with 1% false positive rate
	bf := NewBloomFilter(10000, 0.01)
	fmt.Println("--- Bloom Filter (10000 items, 1% FP rate) ---")
	bf.Stats()

	// Add items
	for i := 0; i < 10000; i++ {
		bf.Add(fmt.Sprintf("user:%d", i))
	}
	fmt.Println("\nAfter inserting 10000 items:")
	bf.Stats()

	// Test: items that were added should always be found (no false negatives)
	falseNegatives := 0
	for i := 0; i < 10000; i++ {
		if !bf.Contains(fmt.Sprintf("user:%d", i)) {
			falseNegatives++
		}
	}
	if falseNegatives > 0 {
		panic(fmt.Sprintf("FAIL: %d false negatives (should be 0)", falseNegatives))
	}
	fmt.Println("\nPASS: zero false negatives (all inserted items found)")

	// Test: items NOT added - measure false positive rate
	falsePositives := 0
	testItems := 10000
	for i := 10000; i < 10000+testItems; i++ {
		if bf.Contains(fmt.Sprintf("user:%d", i)) {
			falsePositives++
		}
	}
	fpRate := float64(falsePositives) / float64(testItems) * 100
	fmt.Printf("False positives: %d/%d (%.2f%%)\n", falsePositives, testItems, fpRate)

	if fpRate < 2.0 { // Should be close to 1%
		fmt.Println("PASS: false positive rate within expected range")
	} else {
		fmt.Printf("WARNING: FP rate %.2f%% higher than expected 1%%\n", fpRate)
	}

	// Demonstrate: definitely not in set
	if !bf.Contains("definitely-not-a-user") {
		fmt.Println("PASS: 'definitely-not-a-user' correctly identified as absent")
	}

	// Comparison: without bloom filter, checking 10K items in a set requires the full set in memory.
	// Bloom filter: ~12KB vs HashSet: ~800KB+ for 10K strings.
	fmt.Printf("\nMemory: Bloom filter uses only %.2f KB for 10K items\n", float64(bf.size)/8/1024)
	fmt.Println("PASS: bloom filter complete")
}
