package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Halo " + name + " Nama Saya " + customer.Name)
}

func main() {
	pelanggan1 := Customer{}
	pelanggan1.Name = "Pai"
	pelanggan1.Address = "Indonesia"
	pelanggan1.Age = 26

	fmt.Println(pelanggan1)
	fmt.Println(pelanggan1.Name)
	fmt.Println(pelanggan1.Address)
	fmt.Println(pelanggan1.Age)

	pelanggan2 := Customer{
		Name:    "Aminu",
		Address: "Indonesia",
		Age:     26,
	}

	fmt.Println(pelanggan2)
	fmt.Println(pelanggan2.Name)
	fmt.Println(pelanggan2.Address)
	fmt.Println(pelanggan2.Age)

	pelanggan1.sayHello("Admin")
	pelanggan2.sayHello("Admin")
}
