package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func print(n int) int {
	if n == 0 {
		return 0
	}
	fmt.Println(n)
	return print(n - 1)
}

func TOHUtil(n int, from, to, temp string) {
	if n == 1 {
		fmt.Println("Move disk", n, "from peg", from, "to peg", to)
		return
	}

	TOHUtil(n-1, from, temp, to)

	fmt.Println("Move disk", n, "from peg", from, "to peg", to)

	TOHUtil(n-1, temp, to, from)
}

func TowersOfHanoi(n int) {
	TOHUtil(n, "A", "C", "B")
}

func isSorted(A []int) bool {
	n := len(A)
	if n == 1 {
		return true
	}

	if A[n-1] < A[n-2] {
		return false
	}

	return isSorted(A[:n-1])
}

func printResult(A []int, n int) {
	var i int
	for ; i < n; i++ {
		fmt.Print(A[i])
	}
	fmt.Printf("\n")
}

func generateBinaryStrings(n int, A []int, i int) {
	if i == n {
		printResult(A, n)
		return
	}

	A[i] = 0

	generateBinaryStrings(n, A, i+1)

	A[i] = 1
	generateBinaryStrings(n, A, i+1)
}

func generateK_aryStrings(n int, A []int, i int, k int) {
	if i == n {
		printResult(A, n)
		return
	}
	for j := 0; j < k; j++ {
		A[i] = j
		generateK_aryStrings(n, A, i+1, k)
	}
}

func findConnects(matrix [][]int, M, N, r, c int) int {
	answer := 0
	if r < 0 || c < 0 || r >= M || c >= N {
		answer = 0
	} else if matrix[r][c] == 1 {
		matrix[r][c] = 0
		answer = 1 +
			findConnects(matrix, M, N, r-1, c) +
			findConnects(matrix, M, N, r+1, c) +
			findConnects(matrix, M, N, r, c-1) +
			findConnects(matrix, M, N, r, c+1) +
			findConnects(matrix, M, N, r-1, c-1) +
			findConnects(matrix, M, N, r-1, c+1) +
			findConnects(matrix, M, N, r+1, c-1) +
			findConnects(matrix, M, N, r+1, c+1)
	}
	return answer
}

func findMaxConnects(matrix [][]int, M, N int) int {
	maxConnects := 0
	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			maxConnects = max(maxConnects, findConnects(matrix, M, N, r, c))
		}
	}
	return maxConnects
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	matrix := make([][]int, M)
	for i := 0; i < M; i++ {
		matrix[i] = make([]int, N)
		for j := 0; j < N; j++ {
			scanner.Scan()
			v, _ := strconv.Atoi(scanner.Text())
			matrix[i][j] = v
		}
	}

	maxConnects := 0
	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			maxConnects = max(maxConnects, findConnects(matrix, M, N, r, c))
		}
	}
	fmt.Println(findMaxConnects(matrix, M, N))

	/*
		fmt.Println(print(5))

		TowersOfHanoi(3)

		A := []int{10, 20, 23, 23, 45, 78, 88}
		fmt.Println(isSorted(A))

		A = []int{10, 20, 3, 23, 45, 78, 88}
		fmt.Println(isSorted(A))
	*/
}
