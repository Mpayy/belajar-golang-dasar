package main

import "fmt"

// func namaFunction() (typeDataReturn)
func getNameAge() (string, int) {
	// return dataValueReturn
	return "Achmad Rifaih", 26
}

func main() {
	// ini kalau butuh semua variablenya
	name, age := getNameAge()
	fmt.Println("Nama saya", name, "dan Saya Berumur", age, "tahun")
	// ini kalau mau menghiraukan salah satu valuenya, gunakan underscore ( _ )
	name1, _ := getNameAge()
	fmt.Println("Nama saya", name1)
}
