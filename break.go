package main

import "fmt"

func main() {
	// gunakan break kalau ingin keluar dari perulangan
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke ", i)
	}
}
