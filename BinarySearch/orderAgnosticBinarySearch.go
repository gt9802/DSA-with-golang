// it will check if array is sorted in ascending or descending and then do search
package main

import "fmt"

func main() {
	arr1 := []int{40, 38, 33, 28, 25, 18, 15, 10, 5, 2}
	arr2 := []int{2, 5, 6, 8, 12, 18, 22, 34, 56, 60}
	arr3 := []int{-18, -10, -1, 4, 7, 8, 12, 30}

	target1 := 15
	target2 := 6
	target3 := -10

	fmt.Println(orderAgnosticBinarySearch(arr1, target1))
	fmt.Println(orderAgnosticBinarySearch(arr2, target2))
	fmt.Println(orderAgnosticBinarySearch(arr3, target3))
}

func orderAgnosticBinarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	//checking array is asc or desc
	isAsc := arr[start] < arr[end]

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] == target {
			return mid
		}
		if isAsc {
			if target < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target > arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}

	}
	return -1
}
