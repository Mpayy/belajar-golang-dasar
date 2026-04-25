package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddress(address *Address) {
	address.City = "New York"
}

func main() {
	// Bisa begini juga
	//address := &Address{}
	//ChangeAddress(address)

	address := Address{}
	ChangeAddress(&address)

	fmt.Println(address)
}
