package main

import "fmt"

func main() {
	var names [2]string // 2 itu adalah batas maksimal array dan string tipe data yang mau kita isi kedalam array

	names[0] = "Achmad"
	names[1] = "Rifaih"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])

	names[1] = "mpayy"
	fmt.Println(names)

	// membuat array langsung
	var values = [3]int{
		90,
		80,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2]) // kalau ga diisi nilainya default int = 0 string = ""

	// Function Array
	// len(array) mendapatkan panjang array, walau pun kosong kalau batas tetap di isi dengan nilai defaultnya
	// array[index] mendapatkan data di posisi index
	// array[index] = value mengubah data di posisi index
	// di golang tidak bisa menghapus index pada array, paling cuma membuat datanya default

	var numbers = [...]int{ // bisa menggunakan [...] kalau kita tidak mau memberikan batas pada array di awal
		1,
		2,
		3,
		4,
		5,
		6,
	}

	fmt.Println(numbers)
	numbers[4] = 0
	fmt.Println(numbers)
}
