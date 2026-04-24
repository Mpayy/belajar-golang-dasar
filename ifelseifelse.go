package main

import "fmt"

func main() {
	name := "Achmad Rifaih" // didalam pengecekan case sensitif

	if name == "pai" {
		fmt.Println("Halo, pai!!")
	} else if name == "aminu" {
		fmt.Println("Halo, Mantan Terindah!!")
	} else if name == "mita" {
		fmt.Println("Halo, Mantan Terburuk!!")
	} else {
		fmt.Println("Lau siapa?")
	}

	// if short statement
	if lenght := len(name); lenght > 10 {
		fmt.Println("Nama Kepanjangan!!")
	} else if len(name) > 5 {
		fmt.Println("Nama Lumayan Panjang!!")
	} else {
		fmt.Println("Oke Nama Cukup!!")
	}
}
