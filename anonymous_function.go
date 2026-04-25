package main

import "fmt"

type Blacklist func(string) bool

func Registration(name string, blocked Blacklist) {
	if blocked(name) {
		fmt.Printf("You are is already Blocked " + name + "\n")
	} else {
		fmt.Printf("Welcome")
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Mita"
	}
	Registration("Mita", blacklist)

	Registration("Aminu", func(name string) bool {
		return name == "Mita"
	})

}
