package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	//Error
	var number int
	var err error

	number, err = strconv.Atoi("123GHyzxc")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	number, err = strconv.Atoi("157")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}
	println("================================")

	//Error (Custom error), Panic, Recover
	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		// fmt.Println(err.Error())
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Apllication running perfectly")
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("Password has to have ore than 4 characters")
	}

	return "Valid Password", nil
}
