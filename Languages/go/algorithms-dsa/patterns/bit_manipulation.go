package main

import "fmt"

/*
  Pattern: Bit Manipulation
  Templates: Single Number, Power of Two, Counting Bits, Subsets via Bitmask
*/

// singleNumber - every element appears twice except one. XOR cancels pairs.
func singleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}

// isPowerOfTwo - a power of two has exactly one bit set
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// countBits - returns number of 1-bits for each number 0..n
func countBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + (i & 1)
	}
	return dp
}

// hammingDistance - number of positions where bits differ
func hammingDistance(x, y int) int {
	xor := x ^ y
	count := 0
	for xor > 0 {
		count += xor & 1
		xor >>= 1
	}
	return count
}

// reverseBits - reverse the bits of a 32-bit unsigned integer
func reverseBits(n uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result = (result << 1) | (n & 1)
		n >>= 1
	}
	return result
}

// subsetsUsingBitmask - generate all subsets using bit manipulation
func subsetsUsingBitmask(nums []int) [][]int {
	n := len(nums)
	total := 1 << n
	var result [][]int

	for mask := 0; mask < total; mask++ {
		var subset []int
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				subset = append(subset, nums[i])
			}
		}
		result = append(result, subset)
	}
	return result
}

// missingNumber - find the missing number in [0..n]
// XOR of indices and values leaves only the missing number.
func missingNumber(nums []int) int {
	n := len(nums)
	result := n
	for i := 0; i < n; i++ {
		result ^= i ^ nums[i]
	}
	return result
}

// swapWithoutTemp - swap two integers using XOR
func swapWithoutTemp(a, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}

func main() {
	// Test single number
	if singleNumber([]int{2, 2, 1}) != 1 {
		panic("FAIL: singleNumber")
	}
	if singleNumber([]int{4, 1, 2, 1, 2}) != 4 {
		panic("FAIL: singleNumber 2")
	}
	fmt.Println("PASS: singleNumber")

	// Test power of two
	powers := []int{1, 2, 4, 8, 16, 1024}
	for _, p := range powers {
		if !isPowerOfTwo(p) {
			panic(fmt.Sprintf("FAIL: %d should be power of 2", p))
		}
	}
	nonPowers := []int{0, 3, 5, 6, 7, 10}
	for _, p := range nonPowers {
		if isPowerOfTwo(p) {
			panic(fmt.Sprintf("FAIL: %d should NOT be power of 2", p))
		}
	}
	fmt.Println("PASS: isPowerOfTwo")

	// Test count bits
	bits := countBits(5)
	expected := []int{0, 1, 1, 2, 1, 2}
	for i := range expected {
		if bits[i] != expected[i] {
			panic(fmt.Sprintf("FAIL: countBits at %d", i))
		}
	}
	fmt.Println("PASS: countBits:", bits)

	// Test hamming distance
	if hammingDistance(1, 4) != 2 { // 001 vs 100
		panic("FAIL: hamming(1,4)")
	}
	fmt.Println("PASS: hammingDistance")

	// Test reverse bits
	reversed := reverseBits(0b00000000000000000000000000001011)
	if reversed != 0b11010000000000000000000000000000 {
		panic(fmt.Sprintf("FAIL: reverseBits got %032b", reversed))
	}
	fmt.Println("PASS: reverseBits")

	// Test subsets
	subsets := subsetsUsingBitmask([]int{1, 2, 3})
	if len(subsets) != 8 { // 2^3 = 8
		panic(fmt.Sprintf("FAIL: expected 8 subsets, got %d", len(subsets)))
	}
	fmt.Println("PASS: subsetsUsingBitmask, count:", len(subsets))

	// Test missing number
	if missingNumber([]int{3, 0, 1}) != 2 {
		panic("FAIL: missingNumber")
	}
	if missingNumber([]int{0, 1}) != 2 {
		panic("FAIL: missingNumber 2")
	}
	fmt.Println("PASS: missingNumber")

	// Test swap
	a, b := swapWithoutTemp(5, 10)
	if a != 10 || b != 5 {
		panic("FAIL: swap")
	}
	fmt.Println("PASS: swapWithoutTemp")
}
