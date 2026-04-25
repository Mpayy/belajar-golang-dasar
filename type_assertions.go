package main

import "fmt"

func random() any {
	return "Ok" // -> ini akan masuk ke case string
	//return 1 -> ini akan masuk ke case int
	//return true -> ini akan masuk ke default
}

func main() {
	//result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)

	// Kalau salah mengkonversi sebuah data, akan terjadi panic
	//resultInteger := result.(int)
	//fmt.Println(resultInteger)

	// cara ini akan lebih aman, dan tidak menyebabkan panic
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown", value)
	}
}
