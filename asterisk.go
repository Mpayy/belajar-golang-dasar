package main

import "fmt"

type Address struct {
	Kota, Provinsi, Negara string
}

func main() {
	address := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address1 := &address

	// address1 adalah pointer yang menunjuk ke variabel `address`

	// Kalau mengubah field lewat pointer, data asli ikut berubah
	address1.Kota = "Jakarta Timur"

	// Dengan operator * (dereference), kita tidak mengubah field,
	// tapi mengganti SELURUH isi struct yang ditunjuk oleh pointer.
	// Artinya nilai `address` ikut terganti total.
	*address1 = Address{"Subang", "Bandung", "Indonesia"}

	fmt.Println(address)
	fmt.Println(address1)

}
