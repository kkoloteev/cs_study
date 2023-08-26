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
	selectionSort(baseSlice)
	fmt.Println(baseSlice)
}

func selectionSort(baseSlice []int) {
	n := len(baseSlice)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if baseSlice[j] < baseSlice[min] {
				min = j
			}
		}
		baseSlice[i], baseSlice[min] = baseSlice[min], baseSlice[i]
	}
}
