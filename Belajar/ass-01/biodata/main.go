package main

//Melakukan import package otomatis
import (
	"fmt"
	"os"
	"strconv"
)

type Anggota struct {
	//Struct Anggota
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
	no_absen  string
}

func main() {
	//Memasukkan data arg serta merubah menjadi int
	var argsRaw = os.Args
	var args = argsRaw[1]
	angka, err := strconv.Atoi(args)
	//untuk mengubah error serta dapat menampung data yang tidak ada
	_ = err
	anggota1(angka)

}

// func anggota1 yang didalamnya parameter absensi int
func anggota1(absensi int) {
	//Penamaan variable annggotta
	var anggotta = []Anggota{
		//Memanggil struct anggota
		{nama: "Jojon", alamat: "Jl.Sudarso", pekerjaan: "Mahasiswa", alasan: "Ingin Mencoba Belajar Hal Yang Baru", no_absen: "1"},
		{nama: "Amar", alamat: "Jl.Sudarso", pekerjaan: "Mahasiswa", alasan: "Ingin Mencoba Belajar Hal Yang Baru", no_absen: "2"},
		{nama: "Miracle", alamat: "Jl.Sudarso", pekerjaan: "Mahasiswa", alasan: "Ingin Mencoba Belajar Hal Yang Baru", no_absen: "3"},
		{nama: "Ojan", alamat: "Jl.Sudarso", pekerjaan: "Mahasiswa", alasan: "Ingin Mencoba Belajar Hal Yang Baru", no_absen: "4"},
	}

	//Menampilkan data, jika data yang diinput kurang dari sama dengan 4 maka akan menampilkan data di atas
	if absensi <= 4 {
		fmt.Println("Anggota =>")
		fmt.Println("Nama      :", anggotta[absensi-1].nama)
		fmt.Println("Alamat    :", anggotta[absensi-1].alamat)
		fmt.Println("Pekerjaan :", anggotta[absensi-1].pekerjaan)
		fmt.Println("Alasan    :", anggotta[absensi-1].alasan)
		fmt.Println("No Absen  :", anggotta[absensi-1].no_absen)

		//Jika data yang diinputkan lebih dari 4 maka akan menampilkan data tidak ada
	} else {
		fmt.Println("Data Tidak Ada")
	}
}
