package main

import "fmt"

/*
  Pattern: Prefix Sums
  Templates: Range Sum Query, Subarray Sum Equals K, Product of Array Except Self
*/

// PrefixSum array - precompute for O(1) range sum queries
type PrefixSumArray struct {
	prefix []int
}

func NewPrefixSumArray(nums []int) *PrefixSumArray {
	prefix := make([]int, len(nums)+1)
	for i, num := range nums {
		prefix[i+1] = prefix[i] + num
	}
	return &PrefixSumArray{prefix: prefix}
}

// RangeSum returns sum of nums[left..right] inclusive, O(1)
func (p *PrefixSumArray) RangeSum(left, right int) int {
	return p.prefix[right+1] - p.prefix[left]
}

// subarraySum - count subarrays with sum equal to k
// Uses prefix sum + hashmap for O(n) time.
func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	prefixCount := map[int]int{0: 1}

	for _, num := range nums {
		sum += num
		// If (sum - k) was seen before, those positions form valid subarrays
		if c, ok := prefixCount[sum-k]; ok {
			count += c
		}
		prefixCount[sum]++
	}
	return count
}

// productExceptSelf - product of array except self without division
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Left products
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Right products multiplied in
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}
	return result
}

// pivotIndex - find index where left sum == right sum
func pivotIndex(nums []int) int {
	totalSum := 0
	for _, n := range nums {
		totalSum += n
	}
	leftSum := 0
	for i, n := range nums {
		if leftSum == totalSum-leftSum-n {
			return i
		}
		leftSum += n
	}
	return -1
}

// maxSubarraySum - Kadane's algorithm (technically prefix sum variant)
func maxSubarraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if currentSum+nums[i] > nums[i] {
			currentSum += nums[i]
		} else {
			currentSum = nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

// 2D prefix sum for matrix region queries
type PrefixSum2D struct {
	prefix [][]int
}

func NewPrefixSum2D(matrix [][]int) *PrefixSum2D {
	if len(matrix) == 0 {
		return &PrefixSum2D{}
	}
	rows, cols := len(matrix), len(matrix[0])
	prefix := make([][]int, rows+1)
	for i := range prefix {
		prefix[i] = make([]int, cols+1)
	}
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			prefix[i][j] = matrix[i-1][j-1] + prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1]
		}
	}
	return &PrefixSum2D{prefix: prefix}
}

func (p *PrefixSum2D) RegionSum(r1, c1, r2, c2 int) int {
	return p.prefix[r2+1][c2+1] - p.prefix[r1][c2+1] - p.prefix[r2+1][c1] + p.prefix[r1][c1]
}

func main() {
	// Test range sum query
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	psa := NewPrefixSumArray(nums)

	if psa.RangeSum(0, 4) != 15 { // 1+2+3+4+5
		panic("FAIL: RangeSum(0,4)")
	}
	if psa.RangeSum(3, 7) != 30 { // 4+5+6+7+8
		panic("FAIL: RangeSum(3,7)")
	}
	fmt.Println("PASS: PrefixSumArray RangeSum")

	// Test subarray sum equals k
	count := subarraySum([]int{1, 1, 1}, 2)
	if count != 2 {
		panic(fmt.Sprintf("FAIL: subarraySum expected 2, got %d", count))
	}
	count = subarraySum([]int{1, 2, 3}, 3)
	if count != 2 { // [1,2] and [3]
		panic(fmt.Sprintf("FAIL: subarraySum expected 2, got %d", count))
	}
	fmt.Println("PASS: subarraySum")

	// Test product except self
	prod := productExceptSelf([]int{1, 2, 3, 4})
	expected := []int{24, 12, 8, 6}
	for i := range prod {
		if prod[i] != expected[i] {
			panic(fmt.Sprintf("FAIL: productExceptSelf at %d", i))
		}
	}
	fmt.Println("PASS: productExceptSelf:", prod)

	// Test pivot index
	pivot := pivotIndex([]int{1, 7, 3, 6, 5, 6})
	if pivot != 3 {
		panic(fmt.Sprintf("FAIL: pivotIndex expected 3, got %d", pivot))
	}
	fmt.Println("PASS: pivotIndex =", pivot)

	// Test Kadane's
	maxSub := maxSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	if maxSub != 6 { // [4, -1, 2, 1]
		panic(fmt.Sprintf("FAIL: maxSubarray expected 6, got %d", maxSub))
	}
	fmt.Println("PASS: maxSubarraySum =", maxSub)

	// Test 2D prefix sum
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	ps2d := NewPrefixSum2D(matrix)
	regionSum := ps2d.RegionSum(0, 0, 1, 1) // 1+2+4+5 = 12
	if regionSum != 12 {
		panic(fmt.Sprintf("FAIL: 2D region sum expected 12, got %d", regionSum))
	}
	regionSum2 := ps2d.RegionSum(1, 1, 2, 2) // 5+6+8+9 = 28
	if regionSum2 != 28 {
		panic(fmt.Sprintf("FAIL: 2D region sum expected 28, got %d", regionSum2))
	}
	fmt.Println("PASS: PrefixSum2D")
}
