package main

import "fmt"

func main() {
	// Go-lang tidak akan error jika membuat variable dengan const tapi tidak digunakan
	// constant atau const isi datanya tidak bisa dirubah
	//const firstName string = "Achmad"
	//const lastName = "Rifaih"

	// error karena menggunakan constant
	//firstName = "tidak bisa di ubah"
	//lastName = "tidak bisa di ubah"

	const (
		firstName string = "Achmad"
		lastName         = "Rifaih"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
