package main

import "fmt"

func main() {
	// Switch Expression
	name := "rifaih"

	switch name {
	case "pai":
		fmt.Println("Halo, pai!!")
	case "aminu":
		fmt.Println("Halo, aminu!!")
	case "mita":
		fmt.Println("Halo, mita!!")
	default:
		fmt.Println("Halo, Kamu siapa??")
	}

	// Switch short expression
	switch lenght := len(name); lenght > 10 {
	case true:
		fmt.Println("Nama Kepanjangan!!")
	case false:
		fmt.Println("Nama Cukup")
	}

	// Switch tanpa expression, (lebih baik menggnakan if else saja)
	lenght1 := len(name)
	switch {
	case lenght1 > 10:
		fmt.Println("Nama Kepanjangan!!")
	case lenght1 > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Cukup")
	}
}
