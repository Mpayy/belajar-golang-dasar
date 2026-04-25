package main

import "fmt"

type Address struct {
	Kota, Provinsi, Negara string
}

func main() {
	// Passing by Value
	//address := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	//address1 := address // copy value
	//
	//address1.Kota = "Jakarta Timur"
	//fmt.Println(address) // ini tidak berubah
	//fmt.Println(address1) // ini berubah

	// Passing by Reference(Pointer (&))
	address := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address1 := &address // pointer, akan merubah juga data referencenya

	address1.Kota = "Jakarta Timur"
	fmt.Println(address)  // ini  berubah
	fmt.Println(address1) // ini berubah

}
