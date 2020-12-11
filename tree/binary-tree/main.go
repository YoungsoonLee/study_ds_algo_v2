package main

import (
	"fmt"
	"math"
	"math/rand"
)

type BinaryTreeNode struct {
	left  *BinaryTreeNode
	data  int
	right *BinaryTreeNode
}

func PreOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d", root.data)

	PreOrder(root.left)
	PreOrder(root.right)
}

func PreOrderWalk(root *BinaryTreeNode, ch chan int) {
	if root == nil {
		return
	}

	ch <- root.data

	PreOrderWalk(root.left, ch)
	PreOrderWalk(root.right, ch)
}

func PreOrderWalker(root *BinaryTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		PreOrderWalk(root, ch)
		close(ch)
	}()
	return ch
}

func NewBinaryTree(n, k int) *BinaryTreeNode {
	var root *BinaryTreeNode
	for _, v := range rand.Perm(n) {
		root = insert(root, (1+v)*k)
	}
	return root
}

// !!!!
func insert(root *BinaryTreeNode, v int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{nil, v, nil}
	}

	if v < root.data {
		root.left = insert(root.left, v)
		return root
	}

	root.right = insert(root.right, v)
	return root
}

func InOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	InOrder(root.left)
	fmt.Printf("%d", root.data)
	InOrder(root.right)
}

func InOrderWalk(root *BinaryTreeNode, ch chan int) {
	if root == nil {
		return
	}

	InOrderWalk(root.left, ch)
	ch <- root.data
	InOrderWalk(root.right, ch)
}

func InOrderWalker(root *BinaryTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		InOrderWalk(root, ch)
		close(ch)
	}()
	return ch
}

func PostOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Printf("%d", root.data)
}

func PostOrderWalk(root *BinaryTreeNode, ch chan int) {
	if root == nil {
		return
	}

	PostOrderWalk(root.left, ch)
	PostOrderWalk(root.right, ch)
	ch <- root.data
}

func PostOrderWalker(root *BinaryTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		PostOrderWalk(root, ch)
		close(ch)
	}()
	return ch
}

// BFS
// level order
func LevelOrder(root *BinaryTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*BinaryTreeNode{root}

	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			level = append(level, node.data)
			queue = queue[1:]

			if node.left != nil {
				queue = append(queue, node.left)
			}

			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		result = append(result, level)
	}
	return result
}

func findMax(root *BinaryTreeNode) int {
	max := math.MinInt32
	if root != nil {
		root_val := root.data
		left := findMax(root.left)
		right := findMax(root.right)

		if left > right {
			max = left
		} else {
			max = right
		}

		if root_val > max {
			max = root_val
		}
	}
	return max
}

func findMax2(root *BinaryTreeNode) int {
	max := math.MinInt32
	if root == nil {
		return max
	}

	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			if node.data > max {
				max = node.data
			}
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return max
}

func find(root *BinaryTreeNode, data int) *BinaryTreeNode {
	if root == nil {
		return root
	} else {
		if data == root.data {
			return root
		} else {
			temp := find(root.left, data)
			if temp != nil {
				return temp
			} else {
				return find(root.right, data)
			}
		}
	}
}

// !!!
func Insert(root *BinaryTreeNode, v int) *BinaryTreeNode {
	newNode := &BinaryTreeNode{nil, v, nil}
	if root == nil {
		return newNode
	}
	if root.left == nil {
		root.left = Insert(root.left, v)
	} else if root.right == nil {
		root.right = Insert(root.right, v)
	}
	return root
}

func Insert2(root *BinaryTreeNode, v int) *BinaryTreeNode {
	newNode := &BinaryTreeNode{nil, v, nil}
	if root == nil {
		return newNode
	}

	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			} else {
				node.left = newNode
				return root
			}
			if node.right != nil {
				queue = append(queue, node.right)
			} else {
				node.right = newNode
				return root
			}
		}
	}
	return root
}

func Size(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Size(root.left) + Size(root.right)
}

func Size2(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	var result int
	queue := []*BinaryTreeNode{root}

	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			result++
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}

	return result
}

func LevelOrderBottomUp(root *BinaryTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*BinaryTreeNode{root}
	stack := NewStack(1)

	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			level = append(level, node.data)
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.left)
			}
		}
		stack.Push(level)
	}

	for !stack.IsEmpty() {
		result = append(result, stack.Pop().([]int)) // !!!!
	}
	return result
}

func DeleteTree(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	}
	// first delete both subtrees
	root.left = DeleteTree(root.left)
	root.right = DeleteTree(root.right)
	// Delete current node only affter deleting subtree
	root = nil
	return root
}

func Height(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	} else {
		leftHeight := Height(root.left)
		rightHeight := Height(root.right)

		if leftHeight > rightHeight {
			return leftHeight + 1
		} else {
			return rightHeight + 1
		}
	}
}

func Height2(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*BinaryTreeNode{root}
	count := 0
	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.data)
			if node.left != nil {
				queue = append(queue, node.left)
			}

			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		count++
	}
	return count
}

func LeavesCount(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	queue := []*BinaryTreeNode{root}

	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.left != nil {
				queue = append(queue, node.left)
			}

			if node.right != nil {
				queue = append(queue, node.right)
			}

			if node.left == nil && node.right == nil {
				count++
			}
		}
	}

	return count
}

func FullNodesCount(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	queue := []*BinaryTreeNode{root}

	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.left != nil && node.right != nil {
				count++
			}

			if node.left != nil {
				queue = append(queue, node.left)
			}

			if node.right != nil {
				queue = append(queue, node.right)
			}

		}
	}
	return count
}

func HalfNodesCount(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	queue := []*BinaryTreeNode{root}

	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]

			if (node.left != nil && node.right == nil) || (node.right != nil && node.left == nil) {
				count++
			}

			if node.left != nil {
				queue = append(queue, node.left)
			}

			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return count
}

func CompareStructures(root1, root2 *BinaryTreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	return CompareStructures(root1.left, root2.left) && CompareStructures(root1.right, root2.right)
}

func main() {
	t1 := NewBinaryTree(10, 1)

	PreOrder(t1)
	fmt.Println()

	c := PreOrderWalker(t1)
	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("%d", v)
	}

	fmt.Println()

	InOrder(t1)
	fmt.Println()

	c = InOrderWalker(t1)
	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("%d", v)
	}

	fmt.Println()
	fmt.Println(LevelOrder(t1))

	fmt.Println(findMax2(t1))

	fmt.Println(Size(t1))
	fmt.Println(Size2(t1))
}
