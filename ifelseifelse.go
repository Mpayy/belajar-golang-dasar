package main

import "fmt"

func main() {
	name := "mita" // didalam pengecekan case sensitif

	if name == "pai" {
		fmt.Println("Halo, pai!!")
	} else if name == "aminu" {
		fmt.Println("Halo, Mantan Terindah!!")
	} else if name == "mita" {
		fmt.Println("Halo, Mantan Terburuk!!")
	} else {
		fmt.Println("Lau siapa?")
	}
}
