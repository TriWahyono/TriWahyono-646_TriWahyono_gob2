package main

import "fmt"

func main() {
	//Belajar MAP
	var person map[string]string //deklarasi
	person = map[string]string{} //inisialisasi

	person["name"] = "Tri"
	person["age"] = "23"
	person["address"] = "Jl.Wahid Hasin"

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])

	fmt.Print("============================\n")

	//Belajar Map (Looping with map)
	var person1 = map[string]string{
		"name":    "Tri",
		"age":     "23",
		"address": "Jl. Wahid",
	}

	for key, value := range person1 {
		fmt.Println(key, ":", value)
	}
	fmt.Print("============================\n")

	//Belajar Map (Deleting value)
	var person2 = map[string]string{
		"name":    "Tri",
		"age":     "23",
		"address": "Jl. Wahid",
	}
	fmt.Println("Before Del:", person2)

	delete(person2, "age")
	fmt.Println("After Del:", person2)

	fmt.Print("============================\n")

	//Belajar Map (Detecting a value)
	var person3 = map[string]string{
		"name":    "Tri",
		"age":     "23",
		"address": "Jl. Wahid",
	}
	value, exist := person3["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value does'nt exist")
	}
	delete(person3, "age")

	value, exist = person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}
	fmt.Print("============================\n")

	//Belajar Map (Combining slice with map)
	var people = []map[string]string{
		{"name": "Tri", "age": "23"},
		{"name": "Miracle", "age": "22"},
		{"name": "Wahyo", "age": "21"},
	}

	for i, person4 := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person4["age"])
	}
}
