package main

import (
	"fmt"
	"strings"
)

func main() {
	profile("AI", "WE", "QWERTY", "QZA", "AAAAA")
}

func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("Hello there!! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
