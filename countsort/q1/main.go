package main

import "fmt"

func main() {
	arr := []int{5, 1, 3, 2, 5, 2, 6, 8, 2, 0, 4, 5, 1, 6, 8, 2, 7, 9, 2, 1, 5, 6, 10}

	var count [11]int
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	fmt.Println("count:", count)

	for i := 1; i < 11; i++ {
		count[i] += count[i-1]
	}

	fmt.Println("count2:", count)

	sorted := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	fmt.Println(sorted)
}
