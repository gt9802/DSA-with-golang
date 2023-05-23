package main

import (
	"math"
)

// func main() {
// 	arr := []int{2, 10, 7, 20, 1}
// 	fmt.Println(minInList(arr))
// }

func minInList(arr []int) int {
	ans := math.MaxInt16
	for i := 0; i < len(arr); i++ {
		if arr[i] < ans {
			ans = arr[i]
		}
	}
	return ans
}
