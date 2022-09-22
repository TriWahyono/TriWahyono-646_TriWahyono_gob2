package main

import "fmt"

func main() {

	//Belajar Kondisi (Temporary Variable)
	var tahun = 2022

	if age := tahun - 1999; age < 17 {
		fmt.Println("Kamu belum dapat membuat SIM")
	} else {
		fmt.Println("Kamu sudah dapat membuat SIM")
	}
	fmt.Print("============================\n")

	//Belajar Kondisi (Switch)
	var hasil = 6

	switch hasil {
	case 10:
		fmt.Println("LULUS")
	case 7:
		fmt.Println("CUKUP")
	case 6:
		fmt.Println("KURANG")
	}
	fmt.Print("============================\n")

	//Belajar Kondisi (Switch With Relation Operators)
	var hasill = 6

	switch {
	case hasill == 8:
		fmt.Println("LULUS")
	case (hasill < 8) && (hasill > 6):
		fmt.Println("CUKUP")
	default:
		{
			fmt.Println("Belajar Lebih Giat Lagi")
			fmt.Println("Pertahankan")
		}
	}
	fmt.Print("============================\n")

	//Belajar Kondisi (fallthrough Keyword)
	var hasilll = 4

	switch {
	case hasilll == 10:
		fmt.Println("LULUS")
	case (hasilll < 8) && (hasilll > 6):
		fmt.Println("CUKUP")
		fallthrough
	case hasilll < 6:
		fmt.Println("Tolong Belajar Lebih Giat Lagi..!!")
	default:
		{
			fmt.Println("Belajar Lebih Giat Lagi")
			fmt.Println("Pertahankan")
		}
	}
	fmt.Print("============================\n")

	//Belajar Kondisi (Nested Conditions)
	var hasiil = 0

	if hasiil > 7 {
		switch hasiil {
		case 10:
			fmt.Println("LULUS")
		default:
			fmt.Println("PERTAHANKAN")
		}
	} else {
		if hasiil == 6 {
			fmt.Println("TINGKATKAN LAGI")
		} else if hasiil == 3 {
			fmt.Println("BELAJAR LEBIH GIAT")
		} else {
			fmt.Println("KAMU PASTI BISA")
			if hasiil == 0 {
				fmt.Println("JANGAN PUTUS ASA")
			}
		}
	}
	fmt.Print("============================\n")

}
