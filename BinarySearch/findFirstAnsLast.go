package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	ans := []int{-1, -1}
	first := search(nums, target, true)
	last := search(nums, target, false)
	ans[0] = first
	ans[1] = last
	fmt.Println(ans)
}
func search(nums []int, target int, findFirstIndex bool) int {
	ans := -1
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2

		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			ans = mid
			if findFirstIndex {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}

	}
	return ans
}
