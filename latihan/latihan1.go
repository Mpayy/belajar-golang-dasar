package main

import (
	"errors"
	"fmt"
)

// Buat function yang menerima slice of int, return nilai terbesar

// func sortFromLarge(slice ...int) {
// 	for i := 0; i < len(slice); i++{
// 		for j := i + 1; j < len(slice); j++{
// 			if slice[i] < slice[j] {
// 				slice[i], slice[j] = slice[j], slice[i]
// 			}
// 		}
// 	}
// 	fmt.Print(slice)
// }

func returnLargeValue (slice ...int) (int, error){
	if len(slice) == 0{
		return 0, errors.New("Slice Tidak Boleh Kosong")
	}

	result := slice[0]
	for i := 1; i < len(slice); i++ {
		 if result < slice[i] {
			result = slice[i]
		 }
	}
	return result, nil
}

func main(){
	angka := []int{1,2,3,4,5,6,7,8,9}
	val, err := returnLargeValue(angka...)
	// sortFromLarge(slice...)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(val)
}