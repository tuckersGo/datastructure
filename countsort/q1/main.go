package main

import "fmt"

func countsort1(arr []int) []int {
	var count [11]int
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	sorted := make([]int, 0, len(arr))
	for i := 0; i < 11; i++ {
		for j := 0; j < count[i]; j++ {
			sorted = append(sorted, i)
		}
	}
	return sorted
}

func countsort2(arr []int) []int {
	var count [11]int
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	for i := 1; i < 11; i++ {
		count[i] += count[i-1]
	}

	sorted := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	return sorted
}

func main() {
	arr := []int{5, 1, 3, 2, 5, 2, 6, 8, 2, 0, 4, 5, 1, 6, 8, 2, 7, 9, 2, 1, 5, 6, 10}

	fmt.Println(countsort1(arr))
	fmt.Println(countsort2(arr))
}
