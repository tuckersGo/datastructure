package sort

import "golang.org/x/exp/constraints"

func BinaryInsertSort[T constraints.Ordered](sorted []T, val T) []T {
	idx := findInsert(sorted, val)
	return append(sorted[:idx], append([]T{val}, sorted[idx:]...)...)
}

func findInsert[T constraints.Ordered](sorted []T, val T) int {
	if len(sorted) == 0 {
		return 0
	}
	mid := len(sorted) / 2
	if sorted[mid] < val {
		return findInsert(sorted[mid+1:], val) + mid + 1
	} else {
		return findInsert(sorted[:mid], val)
	}
}
