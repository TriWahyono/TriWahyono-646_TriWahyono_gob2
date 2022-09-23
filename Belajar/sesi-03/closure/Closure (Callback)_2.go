package main

import "fmt"

type isOddNum1 func(int) bool

func main() {
	//Belajar Closure (Callback) / isOddNum
	var numbers = []int{2, 5, 8, 10, 3, 99, 23}

	var find = findOddNumbers1(numbers, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find)
}
func findOddNumbers1(numbers []int, callback isOddNum1) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
