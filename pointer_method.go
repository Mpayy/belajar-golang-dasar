package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	pai := Man{"Pai"}
	pai.Married()

	fmt.Println(pai.Name)

}
