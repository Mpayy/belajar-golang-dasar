package main

import "fmt"

func main() {
	// map[tipe data keynya]tipe data valuenya
	person := map[string]string{
		"name":    "Achmad Rifaih",
		"age":     "26",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])

	// fmt.Println(person["salah"]) kalau key salah atau tidak ada, maka akan di isi nilai defaultnya

	// Function Map
	// - len(map) Untuk mendapatkan jumlah data di map
	// - map[key] untuk mengambil data di map berdasarkan key
	// - map[key] = value untuk merubah data di map dengan keynya
	// - make(map[TypeKey]TypeValue) untuk membuat map baru
	// - delete(map, key) untuk menghapus data di map dengan key

	books := make(map[string]string)
	books["judul"] = "Tunggu Aku Sukses Nanti"
	books["author"] = "Achmad Rifaih"
	books["harga"] = "100000"

	fmt.Println(books)
	fmt.Println(len(books))
	fmt.Println(books["judul"])
	fmt.Println(books["author"])

	books["author"] = "Galang Rambu Anarki"
	fmt.Println(books)

	delete(books, "harga")
	fmt.Println(books)
}
