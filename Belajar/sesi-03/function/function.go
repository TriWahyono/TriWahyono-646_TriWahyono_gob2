package main

import (
	"fmt"
	"strings"
)

func main() {
	//Belajar Function
	greet("Tri", "Puluhdadi")
}

func greet(name, address string) {
	fmt.Println("Haloo! Namaku adakah", name)
	fmt.Println("Aku tinggal di", address)
	fmt.Print("============================\n")

	//Belajar Function (Return)
	var names = []string{"Tri", "Wah"}
	var printMsg = greets("Halo", names)

	fmt.Println(printMsg)
}
func greets(msg string, names []string) string {
	var joinStr = strings.Join(names, "")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result

}
