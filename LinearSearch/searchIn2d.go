package main

import "math"

// func main() {
// 	arr := [][]int{{1, 2, 3}, {4, 50, 6}, {7, 8}}
// 	target := 4

// 	fmt.Println(search(arr, target))
// 	fmt.Println("maximum value in 2d array is:", maxIn2d(arr))
// }

func search(arr [][]int, target int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}

}

func maxIn2d(arr [][]int) int {
	ans := math.MinInt16
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] > ans {
				ans = arr[i][j]
			}
		}
	}
	return ans
}
