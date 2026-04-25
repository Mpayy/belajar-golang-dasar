package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("Payy")
	if person == nil {
		fmt.Println("Data person masih kosong")
	} else {
		fmt.Println(person["name"])
	}
}
