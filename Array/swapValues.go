package main

import "fmt"

func SwapValuesInArray(arr []int, index1 int, index2 int) {
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp

	fmt.Println(arr)
}
