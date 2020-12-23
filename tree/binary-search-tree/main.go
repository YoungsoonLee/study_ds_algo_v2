package main

import "fmt"

type BSTNode struct {
	data  int
	left  *BSTNode
	right *BSTNode
}

func find(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return nil
	}

	if data < root.data {
		return find(root.left, data)
	} else if data > root.data {
		return find(root.right, data)
	}
	return root
}

func find2(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return nil
	}

	for root != nil {
		if data < root.data {
			root = root.left
		} else if data > root.data {
			root = root.right
		} else {
			return root
		}
	}
	return nil
}

func findMin(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	} else if root.left == nil {
		return root
	} else {
		return findMin(root.left)
	}
}

func findMin2(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	}

	for root.left != nil {
		root = root.left
	}
	return root
}

func Insert(root *BSTNode, v int) *BSTNode {
	if root == nil {
		return &BSTNode{v, nil, nil}
	}

	if v < root.data {
		root.left = Insert(root.left, v)
		return root
	}
	root.right = Insert(root.right, v)
	return root
}

func Delete(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return nil
	}

	if data < root.data {
		root.left = Delete(root.left, data)
	} else if data > root.data {
		root.right = Delete(root.right, data)
	} else {
		if root.right == nil {
			return root.left
		}
		if root.left == nil {
			return root.right
		}

		t := root
		root = findMin(t.right)
		root.right = deleteMin(t.right)
		root.left = t.left
	}
	return root
}

func deleteMin(root *BSTNode) *BSTNode {
	if root.left == nil {
		return root.right
	}
	root.left = deleteMin(root.left)
	return root
}

// inorder
func Walk(root *BSTNode, ch chan int) {
	if root == nil {
		return
	}

	Walk(root.left, ch)
	ch <- root.data
	Walk(root.right, ch)
}

func Walker(root *BSTNode) <-chan int {
	ch := make(chan int)
	go func() {
		Walk(root, ch)
		close(ch)
	}()
	return ch
}

func Compare(t1, t2 *BSTNode) bool {
	c1, c2 := Walker(t1), Walker(t2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}

func LCA(root *BSTNode, a, b int) *BSTNode {
	cur := root
	for {
		switch {
		case a < cur.data && b < cur.data:
			cur = cur.left
		case a > cur.data && b > cur.data:
			cur = cur.right
		default:
			return cur
		}
	}

	return root
}

func IsBST(root *BSTNode) bool {
	if root == nil {
		return true
	}

	if root.left != nil && root.left.data > root.right.data {
		return false
	}

	if root.right != nil && root.right.data < root.left.data {
		return false
	}

	if !IsBST(root.left) || !IsBST(root.right) {
		return false
	}
	return true
}

func IsBST2(root *BSTNode, prev *int) bool {
	if root == nil {
		return true
	}

	if !IsBST2(root.left, prev) {
		return false
	}

	if root.data < *prev {
		return false
	}
	*prev = root.data

	return IsBST2(root.right, prev)
}

func BST2DLL(root *BSTNode) {
	if root == nil || (root.left == nil && root.right == nil) {
		return
	}
	BST2DLL(root.left)
	BST2DLL(root.right)

	currRight := root.right
	root.right = root.left
	root.left = nil
	for root.right != nil {
		root = root.right
	}
	root.right = currRight
}

func SortedArrayToBST(A []int) *BSTNode {
	if A == nil {
		return nil
	}

	return helper(A, 0, len(A)-1)
}

func helper(A []int, low int, high int) *BSTNode {
	if low > high {
		return nil
	}

	mid := low + (high-low)/2
	node := new(BSTNode)
	node.data = A[mid]
	node.left = helper(A, low, mid-1)
	node.right = helper(A, mid+1, high)
	return node
}

func kthSmallest(root *BSTNode, k int) *BSTNode {
	counter := 0
	return helper2(root, k, &counter)
}

func helper2(root *BSTNode, k int, counter *int) *BSTNode {
	if root == nil {
		return nil
	}

	left := helper2(root.left, k, counter)
	if left != nil {
		return left
	}

	*counter++
	if *counter == k {
		return root
	}

	return helper2(root.right, k, counter)
}

func FloorInBST(root *BSTNode, key int) *BSTNode {
	if root == nil {
		return root
	}

	if key > root.data {
		r := FloorInBST(root.right, key)
		if r == nil {
			return root
		} else {
			return r
		}
	} else if key < root.data {
		return FloorInBST(root.left, key)
	} else {
		return root
	}
}

func CeilInBST(root *BSTNode, key int) *BSTNode {
	if root == nil {
		return root
	}

	if root.data == key {
		return root
	} else if root.data < key {
		return CeilInBST(root.right, key)
	} else {
		l := CeilInBST(root.left, key)
		if l != nil {
			return l
		}
	}
	return root
}

func RangePrinter(root *BSTNode, K1, K2 int) {
	if root == nil {
		return
	}

	if root.data >= K1 {
		RangePrinter(root.left, K1, K2)
	}
	if root.data >= K1 && root.data <= K2 {
		fmt.Print(" ", root.data)
	}
	if root.data <= K2 {
		RangePrinter(root.right, K1, K2)
	}
}

func RangePrinter2(root *BSTNode, K1, K2 int) {
	if root == nil {
		return
	}

	var result [][]int
	queue := []*BSTNode{root}

	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			level = append(level, node.data)
			queue = queue[1:]
			if node.data >= K1 && node.data <= K2 {
				fmt.Print(" ", node.data)
			}
			if node.left != nil && node.data >= K1 {
				queue = append(queue, node.left)
			}
			if node.right != nil && node.data <= K2 {
				queue = append(queue, node.right)
			}
		}
		result = append(result, level)
	}
}
