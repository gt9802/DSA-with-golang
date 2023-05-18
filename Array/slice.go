package main

import "fmt"

func Slice() {
	slice := []int{} // underlying array is allocated in heap and reference to this array is stored in stack
	slice = append(slice, 4)
	fmt.Println(slice)

}
