package main

func BubbleSort(A []int) []int {
	n := len(A) - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
			}
		}
	}
	return A
}

func BubbleSort2(A []int) []int {
	var sorted bool
	items := len(A)
	for !sorted {
		sorted = true
		for i := 1; i < items; i++ {
			if A[i-1] > A[i] {
				A[i-1], A[i] = A[i], A[i-1]
				sorted = false
			}
		}
	}
	return A
}

func SelectionSort(A []int) []int {
	var n = len(A)
	for i := 0; i < n; i++ {
		var minIndex = i
		for j := i; i < n; j++ {
			if A[j] < A[minIndex] {
				minIndex = j
			}
		}
		A[i], A[minIndex] = A[minIndex], A[i]
	}
	return A
}

func InsertionSort(A []int) []int {
	n := len(A)
	for i := 1; i <= n-1; i++ {
		j := i
		for j > 0 {
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
			}
			j -= 1
		}
	}
	return A
}

func MergeSort(A []int) []int {
	if len(A) <= 1 {
		return A
	}

	middle := len(A) / 2
	left := MergeSort(A[:middle])
	right := MergeSort(A[middle:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	for i := 0; len(left) > 0 || len(right) > 0; i++ {
		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if len(left) > 0 {
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			result[i] = right[0]
			right = right[1:]
		}
	}
	return result
}
