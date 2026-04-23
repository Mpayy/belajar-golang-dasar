package main

import "fmt"

func main() {
	// NoKTP ini sebagai alias untuk string
	type NoKTP string

	var ktpPai NoKTP = "123456789"

	var contoh = "987654321"
	var contohKtp = NoKTP(contoh)

	fmt.Println(ktpPai)
	fmt.Println(contohKtp)
}
