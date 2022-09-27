package main

import "fmt"

func main() {
	//empty interface
	var randomValues interface{}
	_ = randomValues

	randomValues = "Yogyakarta"
	randomValues = 2022
	randomValues = true
	randomValues = []string{"Sleman", "Tugu"}

	fmt.Println("=======================================")

	//empty interface (type assertion)
	var v interface{}

	v = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}
	fmt.Println("=======================================")

	//Empty interface (Map & Slice with Empty Interface)
	rs := []interface{}{1, "YK", true, 2, "Tri", true}

	rm := map[string]interface{}{
		"Nama":   "Tri",
		"Status": true,
		"Umur":   22,
	}

	_, _ = rs, rm
}
