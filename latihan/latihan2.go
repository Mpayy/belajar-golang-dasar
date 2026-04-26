package main

import "fmt"

// Buat struct Mahasiswa dengan method Perkenalan() yang print nama dan umur

type Mahasiswa struct{
	nama string
	umur int
}

func (m Mahasiswa) Perkenalan(){
	fmt.Println("Perkenalkan nama saya: ", m.nama)
	fmt.Println("Umur Saya: ", m.umur)
}

func main(){
	mahasiswa1 := Mahasiswa{"Achmad Rifaih", 26}
	mahasiswa1.Perkenalan()
}