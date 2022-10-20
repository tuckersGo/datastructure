package sort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(100)
	}

	assert.False(t, IsSorted(arr))
	QuickSort(arr)
	assert.True(t, IsSorted(arr), arr)
	t.Log(arr)
}

func TestMergeSort(t *testing.T) {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(100)
	}

	assert.False(t, IsSorted(arr))
	sorted := MergeSort(arr)
	assert.True(t, IsSorted(sorted), arr)
	t.Log(sorted)
}

func TestInsertSort(t *testing.T) {
	arr := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		arr = BinaryInsertSort(arr, rand.Intn(100))
	}
	assert.True(t, IsSorted(arr), arr)
	t.Log(arr)
}

func BenchmarkQuickSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}

	QuickSort(arr)
}

func BenchmarkMergeSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}

	MergeSort(arr)
}

func BenchmarkInsertSort(b *testing.B) {
	arr := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		arr = BinaryInsertSort(arr, rand.Intn(b.N))
	}
}
