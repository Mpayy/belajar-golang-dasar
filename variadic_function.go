package main

import "fmt"

func sumAll(numbers ...int) int { // numbers ini akan menjadi sebuah slice yang menampung parameter
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// kalau ada kasus ternyata data yang di kirim sudah berbentuk slice
	// maka saat mengirimkan ke function, harus di tambah ( ... ).
	// kalau tidak, nanti akan di anggap satu variable, bukan slice atau array
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total := sumAll(numbers...)
	fmt.Println(total)
}
