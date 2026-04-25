package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Mpayy")
	fmt.Println(result)
	// fmt.Println(helper.version) ini tidak bisa di akses karena nama function atau variablenya diawalin huruf kecil
	fmt.Println(helper.Application) 
	// fmt.Println(helper.sayGoodBye) ini tidak bisa di akses karena nama function atau variablenya diawalin huruf kecil
}
