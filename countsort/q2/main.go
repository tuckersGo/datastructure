package main

import "fmt"

func main() {
	str := "dslkjfdslkfjeqwmkljflksdvciuewksdahjsdh"

	var count [26]int
	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++
	}

	maxCount := 0
	var maxCh byte
	for i := 0; i < 26; i++ {
		if count[i] > maxCount {
			maxCount = count[i]
			maxCh = byte('a' + i)
		}
	}

	fmt.Printf("Max character:%c count:%d", maxCh, maxCount)
}
