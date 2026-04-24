package main

import "fmt"

// func namaFunction(parameter, typeDataParameter)
func getName(firstName string, lastName string) {
	fmt.Println(firstName + " " + lastName)
}

func main() {
	getName("Achmad", "Rifaih")
}
