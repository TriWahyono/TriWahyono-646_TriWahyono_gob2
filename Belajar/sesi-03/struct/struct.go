package main

import (
	"fmt"
	"strings"
)

// Belajar Struct (Giving value to struct)
type Employee struct {
	name     string
	age      int
	division string
}

// Belajar Struct (Giving value to struct)
type Person struct {
	name string
	age  int
}

type Employee1 struct {
	division string
	person   Person
}

// Struct (Anonymous struct)

func main() {
	//Belajar Struct (Giving value to struct)
	var employee Employee

	employee.name = "Tri"
	employee.age = 23
	employee.division = "Golang"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	fmt.Print("============================\n")

	//Belajar Struct (Initializing struct)
	var employee1 = Employee{}
	employee1.age = 23
	employee1.division = "Golang"

	var employee2 = Employee{name: "Tri", age: 23, division: "GOLANG"}

	fmt.Printf("Employee1:  %+v\n", employee1)
	fmt.Printf("Employee2: %+v]n", employee2)

	fmt.Print("============================\n")

	//Belajar Struct (Pointer to a struct)
	var employee11 = Employee{name: "Tri", age: 23, division: "GolangHacktive8"}
	var employee22 *Employee = &employee11

	fmt.Println("Employee11 name:", employee11.name)
	fmt.Println("Employee22 name:", employee22.name)

	fmt.Println(strings.Repeat("#", 20))

	employee22.name = "Tri"

	fmt.Println("Employee11 name:", employee11.name)
	fmt.Println("Employee22 name:", employee22.name)

	fmt.Print("============================\n")

	//Belajar Struct (Embedded struct)
	var employee3 = Employee1{}

	employee3.person.name = "CUIRAS"
	employee3.person.age = 300
	employee.division = "INDOPRIDE"

	fmt.Printf("%+v", employee3)

	fmt.Print("============================\n")

	//Belajar Struct (Anonymous struct)
	var employee4 = struct {
		person   Person
		division string
	}{}
	employee4.person = Person{name: "ARIEL CUIRAS", age: 27}
	employee4.division = "GTA ROLEPLAY INDOPRIDE"

	//ANONYMOUS STRUCT DENGAN PENGISIAN FIELD
	var employee5 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "QWERTY", age: 11},
		division: "bocil",
	}
	fmt.Printf("Employee4 : %+v\n", employee4)
	fmt.Printf("Employee5 : %+v\n", employee5)

	fmt.Print("============================\n")

	//Belajar Struct (Slice of struct)
	var people = []Person{
		{name: "A", age: 20},
		{name: "B", age: 21},
		{name: "C", age: 22},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}

	fmt.Print("============================\n")

	//Belajar Struct (Slice of anonymous struct)
	var employee6 = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "AB", age: 24}, division: "Hacktive8"},
		{person: Person{name: "CD", age: 25}, division: "GOLANG"},
		{person: Person{name: "EF", age: 26}, division: "KOMINFO"},
	}
	for _, v := range employee6 {
		fmt.Printf("%+v\n", v)
	}
}
