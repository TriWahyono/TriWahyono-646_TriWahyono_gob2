package main

import (
	"fmt"
	"strings"
)

func main() {
	//Belajar Array
	// var numbers [4]int
	// numbers = [4]int{1, 2, 3, 4}

	// var strings = [3]string{"Tri", "Wahyono", "Golang"}

	// fmt.Printf("%#v\n", numbers)
	// fmt.Printf("%#v\n", strings)
	// fmt.Print("============================\n")

	// //Belajar Array(Modify Element Through Index)
	// var fruitss = [3]string{"Tri", "3", "Kominfo"}
	// fruitss[0] = "Ya"
	// fruitss[1] = "Tidak"
	// fruitss[2] = "Bukan"

	// fmt.Printf("%#v\n", fruitss)
	// fmt.Print("============================\n")

	//Belajar Array(Loop through elements)
	var fruits = [3]string{"TES", "QWERTY", "AUIEO"}
	//Cara Pertama
	for i, v := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))
	//Cara Kedua
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index:%d, Value:%s\n", i, fruits[i])
	}
	fmt.Print("============================\n")
	// Belajar Array(Multidimensional Array)
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d", value)
		}
		fmt.Println()
	}
}
