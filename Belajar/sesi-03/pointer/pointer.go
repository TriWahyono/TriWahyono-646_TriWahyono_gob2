package main

import (
	"fmt"
	"strings"
)

func main() {
	//Belajar Pointer (Memory address)
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (value) :", firstNumber)
	fmt.Println("firstNumber (memori addres) :", &firstNumber)

	fmt.Println("secondNumber (value) :", *secondNumber)
	fmt.Println("secondNumber (memori address) :", secondNumber)

	fmt.Print("============================\n")

	//Belajar Pointer (Changing value through pointer)
	var firstPerson string = "Tri"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("firstPerson (value) :", *secondNumber)
	fmt.Println("firstPerson (memori address) :", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 38), "\n")

	*secondPerson = "Tri"

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("firstPerson (value) :", *secondNumber)
	fmt.Println("firstPerson (memori address) :", secondPerson)

	fmt.Print("============================\n")

	//Belajar Pointer (Pointer as a parameter)
	var a int = 10

	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)
}
func changeValue(number *int) {
	*number = 20
}
