//https://leetcode.com/problems/find-numbers-with-even-number-of-digits/

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(nums))
}

func findNumbers(nums []int) int {
	ans := 0

	for _, number := range nums {
		if even(number) {
			ans++
		}
	}

	return ans
}

func even(number int) bool {
	numberOfDigits := int(math.Log10(float64(number))) + 1
	return numberOfDigits%2 == 0
}
