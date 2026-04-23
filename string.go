package main

import "fmt"

func main() {
	fmt.Println("Achmad") // spasi, tab, atau enter akan dihitung
	fmt.Println("Rifaih")
	fmt.Println("Achmad Rifaih")

	fmt.Println(len("Achmad")) // spasi, tab, atau enter akan dihitung
	fmt.Println("Rifaih"[0])   // Kalau cara ambilnya seperti ini, hasilnya adalah byte dari huruf tersebut
	fmt.Println("Achmad Rifaih"[8])
}

//PS C:\Users\rifai\Developments\GOLANG\belajar-golang-dasar> go run string.go
//Achmad
//Rifaih
//Achmad Rifaih
//6
//82 ini adalah byte dari A
//105 ini adalah byte dari R
