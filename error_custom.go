package main

import "fmt"

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
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Error Validasi: ", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("Erroe Not Found", notFoundErr.Error())
		} else {
			fmt.Println("Unknow Error", err.Error())
		}
	} else {
		// Success
		fmt.Println("Success Save Data")
	}

}
