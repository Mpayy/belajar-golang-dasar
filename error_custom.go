package main

import (
	"fmt"
	"go/types"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validaton Error"}
	}

	if id != "Pai" {
		return &notFoundError{"Data Not Found"}
	}

	return nil
}

func main() {
	err := SaveData("Pai", nil)

	if err != nil{
		// Terjadi kesalahan
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("Error Validasi: ", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Error Not Found", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Unknow Error", err.Error())
		// }

		switch finalError := err.(type){
		case *validationError:
			fmt.Println("Error Validasi: ", finalError.Error())
		case *notFoundError:
			fmt.Println("Error Not Found: ", finalError.Error())
		default:
			fmt.Println("Unknow Error", finalError.Error())
		}
		
	} else {
		// Success
		fmt.Println("Success Save Data")
	}

}
