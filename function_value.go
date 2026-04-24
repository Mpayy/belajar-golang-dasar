package main

import "fmt"

func sapaMantan(nama string) string {
	return "Halo Mantan " + nama
}

func main() {
	mantanTerindah := sapaMantan
	mantanTerburuk := sapaMantan

	fmt.Println(mantanTerindah("Aminu Fatma Arsyah"))
	fmt.Println(mantanTerburuk("Mita Selvira"))
}
