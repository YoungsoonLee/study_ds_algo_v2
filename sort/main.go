package main

import (
	"fmt"
	"sort"
	"time"
)

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

func QuickSort(A []int) {
	recursionSort(A, 0, len(A)-1)
	fmt.Println(A)
}

func recursionSort(A []int, left int, right int) {
	if left < right {
		pivot := partition(A, left, right)
		recursionSort(A, left, pivot-1)
		recursionSort(A, pivot+1, right)
	}
}

func partition(A []int, left int, right int) int {
	for left < right {
		for left < right && A[left] <= A[right] {
			right--
		}
		if left < right {
			A[left], A[right] = A[right], A[left]
			left++
		}
		for left < right && A[left] <= A[right] {
			left++
		}
		if left < right {
			A[left], A[right] = A[right], A[left]
			right--
		}
	}
	return left
}

func CountingSort(A []int, K int) []int {
	bucketLen := K + 1
	C := make([]int, bucketLen)

	sortedIndex := 0
	length := len(A)
	for i := 0; i < length; i++ {
		C[A[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for C[j] > 0 {
			A[sortedIndex] = j
			sortedIndex += 1
			C[j] -= 1
		}
	}
	return A
}

func subtractTime(time1, time2 time.Time) {
	count := 1
	then := time1.Add(time.Duration(-count) * time.Minute)

	fmt.Println("1 minutes ago:", then)
}

func findLagestElement(A []int) int {
	largestElement := 0
	for i := 0; i < len(A); i++ {
		if A[i] > largestElement {
			largestElement = A[i]
		}
	}
	return largestElement
}

func checkWhoWinsTheElection(A []int) int {
	maxCounter, counter, candidate := 0, 0, A[0]
	for i := 0; i < len(A); i++ {
		candidate = A[i]
		counter = 0
		for j := i + 1; j < len(A); j++ {
			if A[i] == A[j] {
				counter++
			}
		}
		if counter > maxCounter {
			maxCounter = counter
			candidate = A[i]
		}
	}
	return candidate
}

func mostFrequent(A []int) int {
	sort.Ints(A)
	currentCounter, maxCounter, res, n := 1, 1, A[0], len(A)
	for i := 1; i < n; i++ {
		if A[i] == A[i-1] {
			currentCounter++
		} else {
			if currentCounter > maxCounter {
				maxCounter = currentCounter
				res = A[i-1]
			}
			currentCounter = 1
		}
	}
	if currentCounter > maxCounter {
		maxCounter = currentCounter
		res = A[n-1]
	}
	return res
}

func merge(A []int, m int, B []int, n int) {
	i := m + n - 1
	j, k := m-1, n-1

	for j >= 0 && k >= 0 {
		if A[j] > B[k] {
			A[i] = A[j]
			j--
		} else {
			A[i] = B[k]
			k--
		}
		i--
	}
	if k >= 0 {
		copy(A[:k+1], B[:k+1])
	}
}

func main() {
	//a := []int{50, 25, 92, 16, 76, 30, 43, 54, 19}
	//QuickSort(a)

	now := time.Now()

	fmt.Println("now:", now)

	count := 1
	then := now.Add(time.Duration(-count) * time.Minute)
	// if we had fix number of units to subtract, we can use following line instead fo above 2 lines. It does type convertion automatically.
	// then := now.Add(-10 * time.Minute)
	fmt.Println("1 minutes ago:", then)
}
