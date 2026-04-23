package main

import "fmt"

func main() {
	// Conversi Nilai
	var nilai32 int32 = 32768
	var nilai64 = int64(nilai32)

	// ini hasilnya akan menjadi negative -32768 , karena maksimal kapasitas int16 adalah 32767, jadi kelebihan 1
	// jadi karena melebihi kapasitas, maka akan terjadi number overflow yaitu kembali ke belakang,
	// dan nilai paling bawah int16 adalah -32767
	var nilai16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// Konversi String
	var name = "Achmad Rifaih"
	var e = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
