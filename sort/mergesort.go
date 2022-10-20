package sort

import "golang.org/x/exp/constraints"

func MergeSort[T constraints.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge[T constraints.Ordered](left, right []T) []T {
	i := 0
	j := 0
	idx := 0
	rst := make([]T, len(left)+len(right))
	for i < len(left) || j < len(right) {
		var leftMerge bool
		if i >= len(left) {
			leftMerge = false
		} else if j >= len(right) {
			leftMerge = true
		} else {
			leftMerge = left[i] < right[j]
		}

		if leftMerge {
			rst[idx] = left[i]
			i++
		} else {
			rst[idx] = right[j]
			j++
		}
		idx++
	}
	return rst
}
