package main

import "fmt"

// VARIABLE
// - Variable adalah tempat untuk menyimpan data.
// - Variable digunakan  agar kita bisa mengakses data yang sama dimanapun kita mau.
// - Di Go-lang Variable hanya bisa menyimpan tipe data yang sama. jika ingin menyimpan data
//yang berbeda-beda jenis, kita harus membuat beberapa variable.
// - Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikut dengan nama variable
//dan tipe datanya

func main() {
	//var name string
	//
	//name = "Achmad Rifaih"
	//fmt.Println(name)
	//
	//name = "Aminu Fatma"
	//fmt.Println(name)

	// Tipe Data Variable
	// - Saat kita membuat variable, maka kita wajib menyebutkan tipe data tersebut.
	// - Namun jika kita langsung menginisialisasikan data variable nya, maka kita tidak
	//wajib menyebutkan tipe data variable nya.

	// Contoh:
	// var name = "Achmad Rifaih" ini sama saja seperti yang di bawah
	name := "Achmad Rifaih" // (:) hanya boleh di pakai saat pertama kali membuat variable
	fmt.Println(name)

	name = "Aminu Fatma"
	fmt.Println(name)

	var (
		// Go-lang akan error jika membuat variable dengan var tapi tidak digunakan
		firstName = "Achmad"
		lastName  = "Rifaih"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(firstName + " " + lastName)
}
