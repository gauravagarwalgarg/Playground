package main

import "fmt"

// Iterator Pattern: Provides a way to access elements of a collection
// sequentially without exposing its underlying representation.
// Example: Custom collection with multiple traversal strategies.

// Iterator interface
type Iterator[T any] interface {
	HasNext() bool
	Next() T
	Reset()
}

// Collection interface
type Iterable[T any] interface {
	CreateIterator() Iterator[T]
	Size() int
}

// Concrete collection: BinaryTree (for in-order traversal example)
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func NewBinaryTree(root *TreeNode) *BinaryTree {
	return &BinaryTree{root: root}
}

func (bt *BinaryTree) CreateIterator() Iterator[int] {
	return NewInOrderIterator(bt.root)
}

func (bt *BinaryTree) Size() int {
	return countNodes(bt.root)
}

func countNodes(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + countNodes(node.Left) + countNodes(node.Right)
}

// InOrderIterator - iterates BST in sorted order (left → root → right)
type InOrderIterator struct {
	root    *TreeNode
	stack   []*TreeNode
	current *TreeNode
}

func NewInOrderIterator(root *TreeNode) *InOrderIterator {
	it := &InOrderIterator{root: root}
	it.current = root
	return it
}

func (it *InOrderIterator) HasNext() bool {
	return it.current != nil || len(it.stack) > 0
}

func (it *InOrderIterator) Next() int {
	for it.current != nil {
		it.stack = append(it.stack, it.current)
		it.current = it.current.Left
	}
	// Pop from stack
	node := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]
	it.current = node.Right
	return node.Value
}

func (it *InOrderIterator) Reset() {
	it.stack = nil
	it.current = it.root
}

// RangeIterator - iterates integers in a range (simpler example)
type RangeIterator struct {
	start   int
	end     int
	step    int
	current int
}

func NewRangeIterator(start, end, step int) *RangeIterator {
	return &RangeIterator{start: start, end: end, step: step, current: start}
}

func (r *RangeIterator) HasNext() bool { return r.current < r.end }
func (r *RangeIterator) Next() int {
	val := r.current
	r.current += r.step
	return val
}
func (r *RangeIterator) Reset() { r.current = r.start }

func main() {
	// Build a BST:       5
	//                  /   \
	//                 3     7
	//                / \   / \
	//               1   4 6   9
	root := &TreeNode{Value: 5,
		Left: &TreeNode{Value: 3,
			Left:  &TreeNode{Value: 1},
			Right: &TreeNode{Value: 4},
		},
		Right: &TreeNode{Value: 7,
			Left:  &TreeNode{Value: 6},
			Right: &TreeNode{Value: 9},
		},
	}

	tree := NewBinaryTree(root)
	it := tree.CreateIterator()

	fmt.Println("--- In-Order BST Traversal ---")
	var inOrder []int
	for it.HasNext() {
		inOrder = append(inOrder, it.Next())
	}
	fmt.Println(inOrder)

	// Verify sorted order
	for i := 1; i < len(inOrder); i++ {
		if inOrder[i] <= inOrder[i-1] {
			panic("FAIL: not in sorted order")
		}
	}
	expected := []int{1, 3, 4, 5, 6, 7, 9}
	for i, v := range expected {
		if inOrder[i] != v {
			panic(fmt.Sprintf("FAIL: expected %d at index %d, got %d", v, i, inOrder[i]))
		}
	}
	fmt.Println("PASS: in-order traversal correct")

	// Range iterator
	fmt.Println("\n--- Range Iterator (0 to 10, step 2) ---")
	rangeIt := NewRangeIterator(0, 10, 2)
	var rangeVals []int
	for rangeIt.HasNext() {
		rangeVals = append(rangeVals, rangeIt.Next())
	}
	fmt.Println(rangeVals)

	if len(rangeVals) == 5 && rangeVals[0] == 0 && rangeVals[4] == 8 {
		fmt.Println("PASS: range iterator correct")
	} else {
		panic("FAIL: range iterator")
	}

	// Reset and iterate again
	rangeIt.Reset()
	if rangeIt.HasNext() && rangeIt.Next() == 0 {
		fmt.Println("PASS: reset works")
	} else {
		panic("FAIL: reset")
	}
}
