package main

import "fmt"

// Function Type Declaration
type Filter func(string) string
type Cek func(string) bool

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Halo" + filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}
	return name
}

func mantan(name string, cek Cek) {
	cekMantan := cek(name)
	if cekMantan == true {
		fmt.Println(name + " is mantan!")
	} else {
		fmt.Println(name + " not mantan!")
	}
}

func cekName(name string) bool {
	if name == "aminu" {
		return true
	}
	return false
}

func main() {
	sayHelloWithFilter("Anjing", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	mantan("aminu", cekName)
	cek := cekName
	mantan("mita", cek)
}
