// https://leetcode.com/problems/peak-index-in-a-mountain-array/
package main

import "fmt"

func main() {
	arr := []int{0, 2, 1, 0}
	fmt.Println(peakIndex(arr))
}
func peakIndex(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start < end {
		mid := start + (end-start)/2

		if arr[mid] > arr[mid+1] {
			//you are in decreasing part
			//this might be answer but look left
			end = mid
		} else {
			//asc part of array
			start = mid + 1
		}
	}
	return start
}
