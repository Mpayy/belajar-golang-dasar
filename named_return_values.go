package main

import "fmt"

// kalau dengan cara seperti, jika kita tidak butuh semua valuenya, nanti saat di tampung valuenya akan menjadi default
func getComplete() (firstName, lastName string, age int) {
	firstName = "Achmad"
	//lastName = "Rifaih"
	//age = 26

	return firstName, lastName, age
}

func main() {
	a, b, c := getComplete()
	fmt.Println(a, b, c)
}
