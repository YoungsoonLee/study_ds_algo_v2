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
}
