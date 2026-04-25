package main

import "fmt"

type Address struct {
	Kota, Provinsi, Negara string
}

func main() {
	var address *Address = new(Address)
	var address1 *Address = address

	address1.Negara = "Brunei"
	fmt.Println(address)
	fmt.Println(address1)
}
