package main

import "fmt"

/*
  LC 307 - Range Sum Query - Mutable
  Topic: Segment Tree
  Difficulty: Medium
  Time: O(log n) update/query | Space: O(n)
*/

type NumArray struct {
	tree []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, 2*n)
	for i := 0; i < n; i++ {
		tree[i+n] = nums[i]
	}
	for i := n - 1; i > 0; i-- {
		tree[i] = tree[2*i] + tree[2*i+1]
	}
	return NumArray{tree, n}
}

func (na *NumArray) Update(index int, val int) {
	i := index + na.n
	na.tree[i] = val
	for i >>= 1; i > 0; i >>= 1 {
		na.tree[i] = na.tree[2*i] + na.tree[2*i+1]
	}
}

func (na *NumArray) SumRange(left int, right int) int {
	sum := 0
	l, r := left+na.n, right+na.n+1
	for l < r {
		if l&1 == 1 {
			sum += na.tree[l]
			l++
		}
		if r&1 == 1 {
			r--
			sum += na.tree[r]
		}
		l >>= 1
		r >>= 1
	}
	return sum
}

func main() {
	na := Constructor([]int{1, 3, 5})
	if na.SumRange(0, 2) == 9 {
		fmt.Println("PASS: SumRange(0,2) = 9")
	} else {
		fmt.Println("FAIL: SumRange(0,2)")
	}

	na.Update(1, 2)
	if na.SumRange(0, 2) == 8 {
		fmt.Println("PASS: SumRange after update = 8")
	} else {
		fmt.Println("FAIL: SumRange after update")
	}
}
