package main

import "fmt"

func main() {
	// gunakan continue kalau ingin melongkap sesuatu
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}
}
