package singlethreaded

import (
	"math/rand"
	"sort"
	"testing"
)

func TestQuickSortRandomSlice(t *testing.T) {
	rand.Seed(1)
	data := make([]int, 20)
	for i := range data {
		data[i] = rand.Intn(100)
	}
	quickSort(data)
	if !sort.IntsAreSorted(data) {
		t.Fatalf("slice not sorted: %v", data)
	}
}
