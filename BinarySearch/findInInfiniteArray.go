// https://www.geeksforgeeks.org/find-position-element-sorted-array-infinite-numbers/
package main

import "fmt"

func main() {
	arr := []int{3, 5, 7, 9, 10, 90, 100, 130, 140, 160, 170, 180, 185, 200, 201, 203}
	target := 170
	fmt.Println(findRange(arr, target))
}

func findRange(arr []int, target int) int {
	start := 0
	end := 1
	//maximize the range to search by increasing exponentialy
	for target > arr[end] {
		newStart := end + 1
		end = end + (end-start+1)*2
		start = newStart
	}
	return bSearch(arr, target, start, end)
}

func bSearch(arr []int, target int, start int, end int) int {
	for start <= end {
		mid := start + (end-start)/2

		if target > arr[mid] {
			start = mid + 1
		} else if target < arr[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
