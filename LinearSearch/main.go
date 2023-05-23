package main

// func main() {
// 	arr := []int{10, 40, 23, 43, 12}
// 	target := 10
// 	ans := linearSearch(arr, target)

// 	if ans {
// 		fmt.Println("Element found ")
// 	} else {
// 		fmt.Println("Element not found")
// 	}
// }

// search for element in array
// for each element in index, check if target is equal to element, if equal return true
// if no element equals target then return flase
func linearSearch(arr []int, target int) bool {

	if len(arr) == 0 {
		return false
	}

	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false

}
