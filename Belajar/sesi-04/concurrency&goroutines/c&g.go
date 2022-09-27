package main

import (
	"fmt"
	"runtime"
	"time"
)

// func main() {
// 	//Goroutines
// 	go goroutine()
// }

// func goroutine() {
// 	fmt.Println("Hello")
// }

func main() {
	fmt.Println("Main Exection Started")

	go firstProcess(8)
	secondProcess(8)

	fmt.Println("No Of Goroutines:", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)
	fmt.Println("Main Execution Ended")
}

func firstProcess(index int) {
	fmt.Println("First Process Func Started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First Process Func Ended")
}

func secondProcess(index int) {
	fmt.Println("Second Process Func Started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second Process Func Ended")
}
