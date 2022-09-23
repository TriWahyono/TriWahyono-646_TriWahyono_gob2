package main

import (
	"fmt"
	"strings"
)

type isOddNum func(int) bool

func main() {
	//Belajar Closure(Declare closure in variable)
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}
		return result
	}
	var numbers = []int{4, 93, 77, 10, 52, 22, 34}

	fmt.Println(evenNumbers(numbers...))

	fmt.Print("============================\n")

	//Belajar Closure (IIFE)
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; 1 >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak")
	fmt.Println(isPalindrome)

	fmt.Print("============================\n")

	//Belajar Closure (Closure as a return value)
	var studentLists = []string{"YELLOW", "WEEN", "IM", "NOT", "THECLOWN"}

	var find = findStudent(studentLists)

	fmt.Println(find("AIUEO"))
}

func findStudent(students []string) func(string) string {

	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s does'nt exist!!!", s)
		}
		return fmt.Sprintf("We Found %s at position %d", s, position+1)
	}

}
