package multithreaded

import (
	"math/rand"
	"runtime"
	"sort"
	"testing"
)

func TestParallelSortRandomSlice(t *testing.T) {
	rand.Seed(1)
	data := make([]int, 20)
	for i := range data {
		data[i] = rand.Intn(100)
	}
	parallelSort(data, runtime.NumCPU())
	if !sort.IntsAreSorted(data) {
		t.Fatalf("slice not sorted: %v", data)
	}
}
