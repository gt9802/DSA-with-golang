package main

import "fmt"

func OneDimensionalArray() {
	arr := [5]int{} // allocated in stack
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Printf("one dimensional array: %v\n", arr)
}
