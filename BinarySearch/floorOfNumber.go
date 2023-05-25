package main

import "fmt"

func main() {
	arr := []int{2, 3, 5, 9, 14, 16, 18}
	target := 15
	fmt.Println(floorOfNumber(arr, target))
}

func floorOfNumber(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2

		if target > arr[mid] {
			start = mid + 1
		} else if target < arr[mid] {
			end = mid - 1

		} else {
			return arr[mid]
		}
	}
	return arr[end]
}
