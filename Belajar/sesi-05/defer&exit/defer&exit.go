package main

import (
	"fmt"
	"os"
)

func main() {
	//Defer #1
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Halo minaaa")
	fmt.Println("Welcome back to go learning center")
	fmt.Println("============================================")

	//Defer #2
	callDeferFunc()
	fmt.Println("Hallooo Semuanya!!")
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")

	fmt.Println("============================================")

	//Exit
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}
