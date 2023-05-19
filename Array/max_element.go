package main

import "fmt"

// imagine that array is not empty
// if empty assign max with some random value
func FindMaxInArray(arr []int) {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Printf("Maximum element in array is: %v\n", max)
}
