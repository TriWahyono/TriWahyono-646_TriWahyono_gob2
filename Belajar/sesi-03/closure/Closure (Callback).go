package main

import "fmt"

func main() {

	//Belajar Closure (Callback)
	var numbers = []int{2, 5, 8, 10, 3, 99, 23}

	var find = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find)
}
func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}