package main

import "fmt"

func ReverseArray(arr []int) {
	start := 0
	end := len(arr) - 1

	for start < end {
		swapValuesInArray(arr, start, end)
		start++
		end--
	}

	fmt.Println(arr)

}

func swapValuesInArray(arr []int, index1 int, index2 int) {
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp

}
