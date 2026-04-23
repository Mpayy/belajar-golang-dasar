package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b
	fmt.Println(c)

	// augmented argument (+=, -=, *=, /=, %=)
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)

	// unary operator
	//++ keterangan a = a + 1
	//-- keterangan a = a - 1
	//- keterangan Negative
	//+ keterangan Positive
	//! keterangan Boolean kebalikan

	var j = 1
	j++ // j = j + 1
	fmt.Println(j)
	j-- // j = j - 1
	fmt.Println(j)
}
