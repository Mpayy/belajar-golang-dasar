package main

import "fmt"

func main() {
	counter := 1

	for counter <= 5 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("Selesai Mengulang")

	// For dengan Statement
	// - Didalam for, kita bisa menambahkan statement. dimana terdapat 2 statement yang bisa di tambahkan
	// - Init Statement, yaitu statement sebelum for di eksekusi
	// - Post Statement, yaitu statement yang akan selalu di eksekusi di akhir tiap perulangan

	// for Init Statement; Kondisi; Post Statement
	for counter1 := 1; counter1 <= 5; counter1++ {
		fmt.Println("Perulangan statement ke ", counter1)
	}

	// For Range

	// Slice
	names := []string{
		"Achmad",
		"Rifaih",
	}

	// kalau ga butuh indexnya bisa di ganti underscore ( _ )
	for index, value := range names {
		fmt.Println("Index: ", index+1, "Name: ", value)
	}

	// Map
	books := map[string]string{
		"judul":  "Tunggu aku sukses nanati",
		"author": "Achmad Rifaih",
	}

	// kalau ga butuh keynya bisa di ganti underscore ( _ )
	for key, value := range books {
		fmt.Println(key, " = ", value)
	}
}
