package main

import (
	"fmt"
	"math/rand"
)

func main() {
	baseSlice := make([]int, 10, 10)

	for i := 0; i < len(baseSlice); i++ {
		baseSlice[i] = rand.Intn(10)
	}
	fmt.Println(baseSlice)
	baseSlice = bubbleSort(baseSlice)
	fmt.Println(baseSlice)
}

func bubbleSort(baseSlice []int) []int {
	n := len(baseSlice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if baseSlice[j] > baseSlice[j+1] {
				baseSlice[j], baseSlice[j+1] = baseSlice[j+1], baseSlice[j]
			}
		}
	}
	return baseSlice
}
