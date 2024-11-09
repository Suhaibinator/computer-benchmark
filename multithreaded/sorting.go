package multithreaded

import (
	"math/rand"
	"runtime"
	"sort"
	"sync"
)

const sortArraySize = 1_000_000_000 // Fixed array size

// SortingBenchmark performs multithreaded sorting
func SortingBenchmark() {
	data := make([]int, sortArraySize)
	for i := range data {
		data[i] = rand.Int()
	}
	parallelSort(data, runtime.NumCPU()*3)
}

// parallelSort performs a parallel sort
func parallelSort(data []int, maxWorkers int) {
	n := len(data)
	if n <= 10000 || maxWorkers <= 1 {
		sort.Ints(data)
		return
	}

	mid := n / 2
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		parallelSort(data[:mid], maxWorkers/2)
	}()
	go func() {
		defer wg.Done()
		parallelSort(data[mid:], maxWorkers-maxWorkers/2)
	}()
	wg.Wait()

	// Merge the two sorted halves
	merge(data, mid)
}

// merge combines two sorted halves of the data slice
func merge(data []int, mid int) {
	temp := make([]int, len(data))
	i, j, k := 0, mid, 0
	for i < mid && j < len(data) {
		if data[i] <= data[j] {
			temp[k] = data[i]
			i++
		} else {
			temp[k] = data[j]
			j++
		}
		k++
	}
	for i < mid {
		temp[k] = data[i]
		i++
		k++
	}
	for j < len(data) {
		temp[k] = data[j]
		j++
		k++
	}
	copy(data, temp)
}
