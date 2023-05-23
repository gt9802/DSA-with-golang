package main

// func main() {
// 	str := "milan"

// 	char := 'o'
// 	fmt.Println(searchInString(str, char))

// }

// check each rune is equal to passed rune
func searchInString(str string, char rune) bool {
	for _, item := range str {
		if char == rune(item) {
			return true
		}
	}
	return false
}
