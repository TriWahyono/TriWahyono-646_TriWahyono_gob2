package main

import "fmt"

func main() {
	//Belajar Loopings (first way of looping)
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}
	fmt.Print("============================\n")

	//Belajar Loopings (second way of looping)
	var ii = 0
	for ii < 3 {
		fmt.Println("Angka", ii)
		ii++
	}
	fmt.Print("============================\n")

	//Belajar Loopings (third way of looping)
	var i = 0
	for {
		fmt.Println("Angka", i)

		i++
		if i == 3 {
			break
		}
	}
	fmt.Print("============================\n")

	//Belajar Loopings (break and continue keywords)
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}
		fmt.Println("Angka", i)
	}
	fmt.Print("============================\n")

	//Belajar Loopings (Nested Looping)
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	fmt.Print("============================\n")

	//Belajar Loopings (Label)
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}
