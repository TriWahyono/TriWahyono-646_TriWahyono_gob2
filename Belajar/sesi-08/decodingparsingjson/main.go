package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      string `json:"age"`
}

func main() {
	var jsonString = `
	{
		"full_name" : "QWERTY",
		"email" 	: "qwerty@gmail.com",
		"age"		: 22
	}
	`
	//Decoding JSON To Struct
	// var result Employee

	//Decoding JSON To Empty Interface
	var temp interface{}

	//Decoding JSON To Struct
	// var err = json.Unmarshal([]byte(jsonString), &result)

	//Decoding JSON To Empty Interface
	var err = json.Unmarshal([]byte(jsonString), &temp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//Decoding JSON To Empty Interface
	var result = temp.(map[string]interface{})

	//Decoding JSON To Struct
	// fmt.Println("full_name", result.FullName)
	// fmt.Println("email", result.Email)
	// fmt.Println("age", result.Age)

	//Decoding JSON To Empty Interface
	fmt.Println("full_name", result["full_name"])
	fmt.Println("email", result["email"])
	fmt.Println("age", result["age"])
}
