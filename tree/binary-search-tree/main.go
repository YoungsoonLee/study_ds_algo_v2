package main

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
