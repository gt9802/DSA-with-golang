package main

// func main() {
// 	arr := []int{10, 8, 5, 4, 3, 2, 1}
// 	target := 10
// 	fmt.Println(binarySearch(arr, target))
// }

func binarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2

		if target < arr[mid] {
			end = mid - 1
		} else if target > arr[mid] {
			start = mid + 1
		} else {
			return mid
		}

	}
	return -1

}
