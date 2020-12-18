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
