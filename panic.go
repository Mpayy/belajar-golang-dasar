package main

import "fmt"

// Panic
// - Panic function adalah function yang bisa kita gunakan untuk menghentikan program
// - Panic function biasanya dipanggil ketika terjadi panic/error pada saat program kita berjalan
// - Saat panic function di panggil, program akan berhenti. namun defer function akan tetap di eksekusi

// Recover
// - Recover function adalah function yang bisa kit gunakan untuk menangkap data panic
// - Dengan recover proses panic akan terhenti, dan di tangkap oleh recover. sehingga program akan tetap berjalan.
func endApp() {
	fmt.Println("End Application")
	message := recover()
	fmt.Println("Terjadi Panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Application Error!!")
	}
}

// Cara yang salah menggunakan recover, ini tetap akan membuat program terhenti.

//func endApp() {
//	fmt.Println("End Application")
//}

//func runApp(error bool) {
//	defer endApp()
//	if error {
//		panic("Application Error!!")
//	}
//	message := recover()
//	fmt.Println("Terjadi Panic", message)
//}

func main() {
	runApp(true)
	fmt.Println("Melanjutkan Application")
}
