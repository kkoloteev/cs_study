package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for value := 0; value <= 11; value++ {
		if idx := binarySearch(slice, value); idx == -1 {
			fmt.Printf("Value %d is not found!\n", value)
		} else {
			fmt.Printf("Value %d is idx %d\n", value, idx)
		}
	}
}

func binarySearch(slice []int, value int) int {
	leftIdx := 0
	rightIdx := len(slice) - 1

	for leftIdx <= rightIdx {
		midIdx := (leftIdx + rightIdx) / 2
		if slice[midIdx] < value {
			leftIdx = midIdx + 1
		} else if slice[midIdx] > value {
			rightIdx = midIdx - 1
		} else {
			return midIdx
		}
	}
	return -1
}
