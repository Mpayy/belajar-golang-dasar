package main

import (
	"errors"
	"fmt"
)

// Buat function yang return 2 nilai: hasil bagi dan sisa bagi (modulo)

func returnBagiDanModulo (nilai, pembagi int) (int,int, error){
	if pembagi == 0 {
		return 0, 0, errors.New("Pembagi Tidak Boleh 0")
	}

	resultBagi := nilai / pembagi
	resultModulo := nilai % pembagi
	return  resultBagi, resultModulo, nil
}

func main(){
	hBagi, hModulo, err := returnBagiDanModulo(20,10)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Hasil Bagi: ", hBagi)
	fmt.Println("Sisa Bagi: ", hModulo)
}