package main

import (
	"encoding/json"
	"fmt"
)

type Employeee struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `[
		{
			"full_name" : "asdqwerty",
			"email"		: "zxvb@gmail.com",
			"age"		: 22
		},
		{
			"full_name" : "hgjtyu",
			"email"		: "asd@gmail.com",
			"age"		: 223
		}	
	]
`
	var result []Employeee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}

}
