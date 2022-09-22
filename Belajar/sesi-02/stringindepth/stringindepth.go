package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//Balajar Looping Over String (byte-by-byte)
	city := "Yogyakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}
	fmt.Print("============================\n")

	//Belajar Looping Over String (byte-by-byte)
	var city1 []byte = []byte{89, 111, 103, 121, 97, 107, 97, 114, 116, 97}

	fmt.Println(string(city1))
	fmt.Print("============================\n")

	//Belajar Looping Over String (rune-by-rune)
	str1 := "tidur"
	str2 := "sleep"

	// fmt.Printf("str1 byte legth => %d\n", len(str1))
	// fmt.Printf("str1 byte legth => %d\n", len(str2))

	fmt.Printf("str1 byte legth => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str1 byte legth => %d\n", utf8.RuneCountInString(str2))

	fmt.Print("============================\n")

	//Belajar Looping Over String (rune-by-rune)
	str := "IAUE"

	for i, s := range str {
		fmt.Printf("index => %d, string => %s\n", i, string(s))
	}

}
