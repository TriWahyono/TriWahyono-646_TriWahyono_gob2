package main

import (
	"fmt"
	"strings"
)

func main() {
	//Belajar Slice
	var fruits = []string{"apel", "jeruk", "semangka"}
	_ = fruits

	fmt.Print("============================\n")

	//Belajar Slice (make function)
	var fruitss = make([]string, 3)
	_ = fruitss

	fmt.Printf("%#v\n", fruitss)
	fmt.Print("============================\n")

	//Belajar Slice (append function)
	var fruitts = make([]string, 3)
	_ = fruitts

	fruitts[0] = "Rusmango"
	fruitts[1] = "Apel"
	fruitts[2] = "Pisang"
	fmt.Printf("%#v\n", fruits)
	fmt.Print("============================\n")
	var fruittss = make([]string, 3)
	fruittss = append(fruittss, "Rusmango", "Apel", "Pisang")
	fmt.Printf("%#v\n", fruittss)
	fmt.Print("============================\n")

	//Belajar (append function with ellipsis)
	var fruits1 = []string{"A", "B", "C"}
	var fruits2 = []string{"D", "E", "F"}

	fruits1 = append(fruits1, fruits2...)

	fmt.Printf("fruits1")

	fmt.Print("============================\n")

	//Belajar (copy function)
	var fruitss1 = []string{"A", "B", "C"}
	var fruitss2 = []string{"D", "E", "F"}

	nn := copy(fruitss1, fruitss2)

	fmt.Println("Fruits1 =>", fruitss1)
	fmt.Println("Fruits2 =>", fruitss2)
	fmt.Println("Copied Elements =>", nn)

	fmt.Print("============================\n")

	//Belajar (Slicing)
	var fruitts1 = []string{"AI", "BI", "CI", "DI", "EI"}

	var fruitts2 = fruitts1[1:4]
	fmt.Printf("%#v\n", fruitts2)

	var fruitts3 = fruitts1[0:]
	fmt.Printf("%#v\n", fruitts3)

	var fruitts4 = fruitts1[:3]
	fmt.Printf("%#v\n", fruitts4)

	var fruitts5 = fruitts1[:]
	fmt.Printf("%#v\n", fruitts5)

	//Belajar (Combining slicing and append)
	var fruits11 = []string{"AI", "BI", "CI", "DI", "EI"}
	fruits11 = append(fruits11[:3], "WE")

	fmt.Printf("%#v\n", fruits11)

	fmt.Print("============================\n")

	//Belajar  (Backing array)
	var fruitss11 = []string{"AI", "BI", "CI", "DI", "EI"}
	var fruitss22 = fruitss11[2:4]

	fruitss22[0] = "WE"
	fmt.Println("fruits1 =>", fruitss11)
	fmt.Println("fruits2 =>", fruitss22)

	fmt.Print("============================\n")

	//Belajar (Cap function)
	var fruittss11 = []string{"AI", "BI", "CI", "DI", "EI"}

	fmt.Println("Fruits1 cap:", cap(fruittss11)) //4
	fmt.Println("Fruits1 len:", len(fruittss11)) //4

	fmt.Println(strings.Repeat("#", 20))

	var fruittss22 = fruittss11[0:3]

	fmt.Println("Fruints2 cap:", cap(fruittss22)) //4
	fmt.Println("Fruints2 len:", len(fruittss22)) //3

	fmt.Println(strings.Repeat("#", 20))

	var fruittss33 = fruittss11[1:]

	fmt.Println("Fruints2 cap:", cap(fruittss33)) //3
	fmt.Println("Fruints2 len:", len(fruittss33)) //3

	fmt.Print("============================\n")

	//Belajar (Creating a new backing array)
	cars := []string{"FordMustang", "Lambo", "H2R", "Zentorno"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "SKYLINE"
	fmt.Println("cars:", cars)
	fmt.Println("NewCars:", newCars)
	fmt.Print("============================\n")
}
