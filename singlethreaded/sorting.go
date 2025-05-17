package singlethreaded

import "math/rand"

// Sorting Benchmark
const DefaultSortSize = 1_000_000

// SortingBenchmark sorts a slice of integers of the given size. When the size
// is non-positive, DefaultSortSize is used.
func SortingBenchmark(size int) {
	if size <= 0 {
		size = DefaultSortSize
	}
	data := make([]int, size)

	// Initialize array with random integers
	for i := range data {
		data[i] = rand.Intn(size)
	}

	// Sort the array
	quickSort(data)
}

// QuickSort Algorithm
func quickSort(a []int) {
	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Partition
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Move pivot to its final place
	a[left], a[right] = a[right], a[left]

	// Recursively sort elements before and after partition
	quickSort(a[:left])
	quickSort(a[left+1:])
}
