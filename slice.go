package main

import "fmt"

func main() {
	names := [...]string{
		"John",
		"Paul",
		"George",
		"Ringo",
		"Mpayy",
		"Aminu",
	}
	// var slice = array[low:high] titik akhir tidak diambil, kalau titik awal di ambil
	slice := names[2:5] // [titik awal yang di tentukan :titik akhir yang di tentukan]
	fmt.Println(slice)

	slice1 := names[:3] // [dari awal : titik akhir yang di tentukan]
	fmt.Println(slice1)

	slice2 := names[3:] // [titik awal yang di tentukan : paling akhir]
	fmt.Println(slice2)

	slice4 := names[:] // kalau ini mengambil semuanya
	fmt.Println(slice4)

	// Function slice
	// len(slice) untuk mendapatkan panjang slice tersebut dari low ke high
	// cap(slice) untuk mendapatkan kapasitas array yang di slice, dari titik awal slice sampai akhir array
	// append(slice, data) untuk membuat slice baru dengan menambah data ke posisi terakhir slice,
	// jika kapasitas array sudah penuh, maka otomatis akan membuat array baru
	// make([]TypeData, length, capacity) untuk membuat slice baru dari array yang beneran baru
	// copy(destination, source) untuk menyalin slice dari source ke destination

	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Juma'at",
		"Sabtu",
		"Minggu",
	}
	fmt.Println(days)

	daysSlice1 := days[5:]
	fmt.Println(daysSlice1)
	daysSlice1[0] = "Sabtu Baru"
	daysSlice1[1] = "MInggu Baru"
	fmt.Println(daysSlice1)
	fmt.Println(days)
	// yang berubah adalah daysnya, karena daySlice1 itu bagian dari array days,
	// jadi slice itu tidak membuat array baru

	daysSlice2 := append(daysSlice1, "Libur Baru")
	// karena array tidak bisa bertambah, maka jika di append logic di intenalnya auto membuat array baru analogi disini daysBaru
	// daysBaru := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Juma'at", "Sabtu Baru", "Minggu Baru", "Libur Baru"}
	daysSlice2[0] = "Sabtu Lama" // saat daysSlice2 dirubah isinya, tidak akan mempengaruhi daysSlice1 dan days, karena dia sudah menjadi array baru
	fmt.Println(daysSlice1)
	fmt.Println(daysSlice2)
	fmt.Println(days)

	// Membuat slice langsung
	newSlice := make([]string, 2, 5) // 2 itu panjangnya, 5 itu maksimal kapasitasnya
	newSlice[0] = "Achmad"
	newSlice[1] = "Rifaih"
	//newSlice[2] = "Aminu" ini akan error, untuk menambahnya harus menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice)) // mengetahui panjang
	fmt.Println(cap(newSlice)) // mengetahui kapasitas maxnya

	newSlice1 := append(newSlice, "Aminu")
	fmt.Println(newSlice1)
	fmt.Println(len(newSlice1))
	fmt.Println(cap(newSlice1))

	newSlice1[0] = "Mita" // ini akan mengubah array newSlice juga karena belum max kapasitasnya kalau memang yang di rubah posisinya masih di array lama
	fmt.Println(newSlice1)
	fmt.Println(newSlice)

	// copy slice
	fromSlice := days[:]                                      // ambil semua data arraynya
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // siapin tempat nampunngnya dan harus pas
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Hati-hati jika membuat array atau slice, mudah tertukar jika dengan pembuatan langsung
	iniArray := [...]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}
	// perbedaanya hanya pada [...] atau jumlah kapasitas array
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	// Jadi pilih array atau slice?
	// Kata PZN kalau kita buat aplikasi menggunakan Golang, kebanyakan kita akan menggunakan slice, jarang array
}
