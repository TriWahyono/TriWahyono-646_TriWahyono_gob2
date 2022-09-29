package main

import (
	"errors"
	"fmt"

	"github.com/howeyc/gopass"
)

func main() {
	//variabel string
	var username string
	//inputan username
	fmt.Println("Silahkan masukkan Username yang telah terdaftar:")
	fmt.Scanln(&username)
	//inputan pass
	fmt.Println("Silahkan masukkan Password yang telah terdaftar:")
	password, _ := gopass.GetPasswdMasked()

	//validasi user pass
	if valid, err := validPassword(username, password); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valid)
	}
}

// fungsi validasi user pass
func validPassword(username string, password []byte) (string, error) {
	validpass := string(password[:])
	//jika user pass tidak sama akan error berikut
	if username != "qwerty" || validpass != "qwerty" {
		return "", errors.New("Username / Passoword salah, silahkan cek kembali")
	}
	//jika user pass sama akan menampilkan berhasi login
	return "Berhasil Login", nil
}
