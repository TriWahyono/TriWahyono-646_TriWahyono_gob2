package main

import (
	"fmt"
	"os"
	"strconv"
)

type Anggota struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
	no_absen  string
}

func main() {

	var argsRaw = os.Args
	var args = argsRaw[1]
	num, err := strconv.Atoi(args)
	_ = err
	anggota1(num)

}

func anggota1(absensi int) {
	var anggotta = []Anggota{
		{nama: "Jojon", alamat: "Jl.Sudarso", pekerjaan: "Mahasiswa", alasan: "Ingin Mencoba Belajar Hal Yang Baru", no_absen: "1"},
		{nama: "Amar", alamat: "Jl.Sudarso", pekerjaan: "Mahasiswa", alasan: "Ingin Mencoba Belajar Hal Yang Baru", no_absen: "2"},
		{nama: "Miracle", alamat: "Jl.Sudarso", pekerjaan: "Mahasiswa", alasan: "Ingin Mencoba Belajar Hal Yang Baru", no_absen: "3"},
		{nama: "Ojan", alamat: "Jl.Sudarso", pekerjaan: "Mahasiswa", alasan: "Ingin Mencoba Belajar Hal Yang Baru", no_absen: "4"},
	}

	if absensi <= 4 {
		fmt.Println("Anggota =>")
		fmt.Println("Nama      :", anggotta[absensi-1].nama)
		fmt.Println("Alamat    :", anggotta[absensi-1].alamat)
		fmt.Println("Pekerjaan :", anggotta[absensi-1].pekerjaan)
		fmt.Println("Alasan    :", anggotta[absensi-1].alasan)
		fmt.Println("No Absen  :", anggotta[absensi-1].no_absen)

	} else {
		fmt.Println("Data Tidak Ada")
	}
}
