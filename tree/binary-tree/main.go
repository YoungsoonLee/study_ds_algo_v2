package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
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

/*
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
*/

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

func DiameterOfBinaryTree(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	var diameter int
	Diameter(root, &diameter)
	return diameter
}

func Diameter(root *BinaryTreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	leftDepth := Diameter(root.left, diameter)
	rightDepth := Diameter(root.right, diameter)

	if leftDepth+rightDepth > *diameter {
		*diameter = leftDepth + rightDepth
	}

	return max(leftDepth, rightDepth) + 1
}

func Diameter2(root *BinaryTreeNode) int {
	diameter := 0

	var depth func(node *BinaryTreeNode) int
	depth = func(node *BinaryTreeNode) int {
		if node == nil {
			return 0
		}

		leftDepth := depth(node.left)
		rightDepth := depth(node.right)
		diameter = max(diameter, leftDepth+rightDepth)

		return max(leftDepth, rightDepth) + 1
	}
	depth(root)
	return diameter
}

func maxLevelSum(root *BinaryTreeNode) (elements []int, maxSum, level int) {
	if root == nil {
		return elements, maxSum, level
	}

	var result [][]int
	levelNumber := 0
	queue := []*BinaryTreeNode{root}

	for len(queue) > 0 {
		qlen := len(queue)
		var currentLevel []int
		sum := 0
		for i := 0; i < qlen; i++ {
			node := queue[0]
			currentLevel = append(currentLevel, node.data)
			sum += node.data
			queue = queue[1:]

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		if sum > maxSum {
			maxSum = sum
			elements = currentLevel
			level = levelNumber
		}
		result = append(result, currentLevel)
		levelNumber++
	}
	return elements, maxSum, level
}

func BinaryTreePaths(root *BinaryTreeNode) []string {
	result := make([]string, 0)
	paths(root, "", &result)
	return result
}

func paths(root *BinaryTreeNode, prefix string, result *[]string) {
	if root == nil {
		return
	}

	if len(prefix) == 0 {
		prefix += strconv.Itoa(root.data)
	} else {
		prefix += "->" + strconv.Itoa(root.data)
	}

	// leaf
	if root.left == nil && root.right == nil {
		*result = append(*result, prefix+"\n")
		return
	}

	paths(root.left, prefix, result)
	paths(root.right, prefix, result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func HasPathSum(root *BinaryTreeNode, sum int) bool {
	allSums := make([]int, 0)
	getAllSums(root, &allSums, 0)
	for _, val := range allSums {
		if sum == val {
			allSums = []int{}
			return true
		}
	}
	allSums = []int{}
	return false
}

func getAllSums(root *BinaryTreeNode, allSums *[]int, currSum int) {
	if root != nil {
		currSum += root.data
		if root.left == nil && root.right == nil {
			*allSums = append(*allSums, currSum)
		} else {
			getAllSums(root.left, allSums, currSum)
			getAllSums(root.right, allSums, currSum)
		}
	}
}

//
func Sum(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	return (root.data + Sum(root.left) + Sum(root.right))
}

func Sum2(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	var result int
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]
			result += node.data
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

func InvertTree(root *BinaryTreeNode) *BinaryTreeNode {
	if root != nil {
		root.left, root.right = InvertTree(root.right), InvertTree(root.left)
	}
	return root
}

func InvertTree2(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	root.left, root.right = root.right, root.left
	InvertTree2(root.left)
	InvertTree2(root.right)
	return
}

func checkMirror(root1, root2 *BinaryTreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.data != root2.data {
		return false
	}

	return checkMirror(root1.left, root2.right) && checkMirror(root1.right, root2.left)
}

func LCA(root *BinaryTreeNode, a, b int) *BinaryTreeNode {
	if root == nil {
		return root
	}

	if root.data == a || root.data == b {
		return root
	}

	left := LCA(root.left, a, b)
	right := LCA(root.right, a, b)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	} else {
		return right
	}
}

func BuildBinaryTree(preOrder []int, inOrder []int) *BinaryTreeNode {
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}
	inOrderIndex := findArray(inOrder, preOrder[0])

	left := BuildBinaryTree(preOrder[1:inOrderIndex+1], inOrder[:inOrderIndex])
	right := BuildBinaryTree(preOrder[inOrderIndex+1:], inOrder[inOrderIndex+1:])

	return &BinaryTreeNode{
		data:  preOrder[0],
		left:  left,
		right: right,
	}

}

func findArray(A []int, target int) int {
	for i, x := range A {
		if x == target {
			return i
		}
	}
	return -1
}

func ZigzagLevelOrder(root *BinaryTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*BinaryTreeNode{root}
	var res [][]int
	leftToright := false
	for {
		qlen := len(queue)
		if qlen == 0 {
			break
		}

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

		if leftToright {
			reverse(level)
		}
		res = append(res, level)
		leftToright = !leftToright
	}
	return res
}

func reverse(list []int) {
	for i := 0; i < len(list)/2; i++ {
		list[i], list[len(list)-1] = list[len(list)-1], list[i]
	}
}

func VerticalTraversal(root *BinaryTreeNode) [][]int {
	data := [][]int{}
	preorder(root, 0, 0, &data)
	sort.Slice(data, func(i, j int) bool {
		if data[i][0] == data[j][0] {
			if data[i][1] == data[j][1] {
				return data[i][2] < data[j][2]
			}
			return data[i][1] > data[j][1]
		}
		return data[i][0] < data[j][0]
	})
	lastX := data[0][0]
	traversal := [][]int{{}}
	for _, v := range data {
		if v[0] != lastX {
			traversal = append(traversal, []int{})
			lastX = v[0]
		}
		traversal[len(traversal)-1] = append(traversal[len(traversal)-1], v[2])
	}
	return traversal
}

func preorder(node *BinaryTreeNode, x, y int, data *[][]int) {
	if node == nil {
		return
	}
	*data = append(*data, []int{x, y, node.data})
	preorder(node.left, x-1, y-1, data)
	preorder(node.right, x+1, y+1, data)
}

/*
func ConnectSiblings(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	}

	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue := queue[1:]
		if node.left != nil && node.right != nil {
			node.left.ne
		}
	}
}
*/

func MinDepth(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	if root.left == nil && root.right == nil {
		return 1
	}

	if root.left == nil {
		return MinDepth(root.right) + 1
	}

	if root.right == nil {
		return MinDepth(root.left) + 1
	}

	return min(MinDepth(root.left), MinDepth(root.right)) + 1
}

func MinDepth2(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	left, right := MinDepth2(root.left), MinDepth2(root.right)
	if left == 0 || right == 0 {
		return left + right + 1
	}
	return min(left, right) + 1
}

func MinDepth3(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	queue := []*BinaryTreeNode{root}

	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			level = append(level, node.data)
			queue = queue[1:]
			if node.left == nil && node.right == nil {
				return count + 1
			}
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printAncestors(root *BinaryTreeNode, node int) bool {
	if root == nil {
		return false
	}

	if root.data == node {
		return true
	}

	left := printAncestors(root.left, node)
	right := false
	if !left {
		right = printAncestors(root.right, node)
	}

	if left || right {
		fmt.Printf("%d ", root.data)
	}

	return left || right
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

	s := "abcdefg"
	f := 0
	l := len(s) - 1
	for f < l {
		s[f], s[l] = s[l], s[f]
	}

	fmt.Println(s)

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

	fmt.Println(BinaryTreePaths(t1))
}
