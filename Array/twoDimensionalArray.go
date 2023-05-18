package main

import "fmt"

func TwoDimensionalArray() {
	arr := [3][3]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	// printing 2D array
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(arr[i][j])

		}
		fmt.Println()
	}
}
