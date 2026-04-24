package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Halo" + filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}
	return name
}

func main() {
	sayHelloWithFilter("Anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
