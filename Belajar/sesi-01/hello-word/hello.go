package main

import "fmt"

func main() {

	// Belajar Hello
	fmt.Println("Hello")
	fmt.Println("Tri Wahyono")
	fmt.Print("============================\n")
	fmt.Print("Hello\n")
	fmt.Print("============================\n")
	fmt.Println("Tri", "Wahyono", "UAJY")
	fmt.Print("============================\n")
	fmt.Print("Ok", " ", "Ya\n")
	fmt.Print("============================\n")

	//Belajar Variabel
	var NamaLengkap string = "Tri Wahyono"
	var Umur int = 23

	fmt.Println("Nama Saya ==>", NamaLengkap)
	fmt.Println("Umur Saya ==>", Umur)
	fmt.Print("============================\n")

	//Belajar Variabel Without Tipe Data
	var NamaPanggilanAnda = "ADMIN"
	var UmurAnda = true

	fmt.Printf("%T, %T\n", NamaPanggilanAnda, UmurAnda)
	fmt.Print("============================\n")

	//Belajar Short Declaration
	NamaSaya := "Tri"
	UmurSaya := 23

	fmt.Printf("%T, %T\n", NamaSaya, UmurSaya)
	fmt.Print("============================\n")

	//Belajar Multiple Variabel Declaration
	var one, two, three string = "1", "2", "3"
	var Angka1, Angka2, Angka3 string = "1", "2", "3"
	fmt.Println("Test Data Multiple Variabel Declaration", one, two, three)
	fmt.Println("Test Data Multiple Variabel Declaration", Angka1, Angka2, Angka3)
	fmt.Print("============================\n")

	//Belajar Multiple Variabel Declaration Short
	var satu, dua, tiga = "1", true, "Tri"
	var Angkaa1, Angkaa2, Angkaa3 = true, 12.30000, "Wahyono"
	fmt.Println("Test Data Multiple Variabel Declaration Short", satu, dua, tiga)
	fmt.Println("Test Data Multiple Variabel Declaration Short", Angkaa1, Angkaa2, Angkaa3)
	fmt.Print("============================\n")

	//Belajar Underscore Variabel &  Print Usage
	var firstVariable string
	var NamaPanggilanMu, NamaPertamaMu, NamaAkhiranMu, NamaLengkapMu = "Tri", "Tri", "Wahyono", "TriWahyono"
	_, _, _, _, _ = firstVariable, NamaPanggilanMu, NamaPertamaMu, NamaAkhiranMu, NamaLengkapMu
	fmt.Print("============================\n")

	//Belajar Print Usage
	var testVariable string
	var NamaMu, PanggilanMu, LengkapMu = "Tri", "Tri", "TriWahyono"
	_, _, _, _ = testVariable, NamaMu, PanggilanMu, LengkapMu
	fmt.Printf("test underscore variable > %T\n", testVariable)
	fmt.Printf("Hello Nama Saya %s, Nama Panggilanku %s, Dan Nama Lengkapku %s\n", NamaMu, PanggilanMu, LengkapMu)
	fmt.Print("============================\n")

	//Belajar Tipe Data Integer
	var Pertama = 646
	var Kedua = -646

	fmt.Printf("Tipe Data satu : %T\n", Pertama)
	fmt.Printf("Tipe Data satu : %T\n", Kedua)
	fmt.Print("============================\n")

	//Belajar Tipe Data Float
	var AngkaDesimal float32 = 646.646

	fmt.Printf("Angka Desimal : %f\n", AngkaDesimal)
	fmt.Printf("Angka Desimal : %.5f\n", AngkaDesimal)
	fmt.Print("============================\n")

	//Belajar Tipe Data Bool
	var kondisi bool = true
	fmt.Printf("Is It Permitted? %t\n", kondisi)
	fmt.Print("============================\n")

	//Belajar Constant
	const NamaPanjang string = "Tri Wahyono"
	fmt.Printf("Hello %s", NamaPanjang)
	fmt.Print("============================\n")

	//Belajar Operators (Arithmetic Operators)
	var value = 2 + 2*3
	fmt.Println(value)
	fmt.Print("============================\n")

	var valuee = (2 + 2) * 3
	fmt.Println(valuee)
	fmt.Print("============================\n")

	//Belajar Relational Operators
	var KondisiPertama bool = 10 > 12
	var KondisiKedua bool = "Tri" == "tri"
	var KondisiKetiga bool = 646 != 6.46

	fmt.Println("Kondisi Pertama :", KondisiPertama)
	fmt.Println("Kondisi Pertama :", KondisiKedua)
	fmt.Println("Kondisi Pertama :", KondisiKetiga)
	fmt.Print("============================\n")

	//Belajar Logical Operator
	var Benar = true
	var Salah = false

	var BenarDanSalah = Salah && Benar
	fmt.Printf("Salah && Benar \t(%t)\n", BenarDanSalah)

	var BenarAtauSalah = Salah || Benar
	fmt.Printf("Salah || Benar \t(%t)\n", BenarAtauSalah)

	var SalahRev = !Salah
	fmt.Printf("!Salah \t\t(%t)\n", !SalahRev)
	fmt.Print("============================\n")

}
