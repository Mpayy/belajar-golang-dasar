package helper

import "fmt"

var version string = "1.0.0"
var Application string = "golang"

func sayGoodBye (name string) string { 
	return "Good Bye" + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh (){
	sayGoodBye("Mpayy") // tetapi ini masih bisa di akses di package yang sama
	fmt.Println(version) // tetapi ini masih bisa di akses di package yang sama
}