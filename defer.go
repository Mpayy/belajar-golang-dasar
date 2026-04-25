package main

import "fmt"

// Defer
// - Defer function adalah function yang bisa kita jadwalkan untuk di eksekusi setelah sebuah function selesai di eksekusi
// - Defer funtion akan selalu di eksekusi walaupun terjadi error pada function yang memanggilnya
func logging() {
	fmt.Println("Selesai Menjalankan Log")
}

//func runApplication() {
//	fmt.Println("Running application")
// jika menggunakan cara ini, kalau terjadi error di tengah dan kita memanggil logging() di akhir, maka logging() tidak akan di eksekusi
// error
//	logging()
//}

func runApplication() {
	defer logging()
	fmt.Println("Running application")
}

func main() {
	runApplication()
}
