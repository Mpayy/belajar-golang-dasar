package main

import "fmt"

// func namaFunction(parameter typeDataParameter) typeDataReturn
func sapaDia(name string) string {
	hello := "Hello " + name + "!"
	return hello
}

func main() {
	result := sapaDia("Achmad")
	fmt.Println(result)

	fmt.Println(sapaDia("Aminu"))
}
