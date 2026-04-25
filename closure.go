package main

import "fmt"

// Closure = Kemampuan sebuah function untuk berinteraksi dengan data-data di sekitarnya
// Harap dengan bijak menggunakan closure, karena:
// - increment adalah closure yang menangkap (capture) variabel `counter` dari scope luar.
// - Artinya setiap pemanggilan increment akan mengubah `counter` yang sama,
// - bukan membuat salinan baru. Perlu hati-hati karena ini bisa menimbulkan
// - side effect yang tidak terlihat jelas jika dipakai di banyak tempat.
func main() {
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
