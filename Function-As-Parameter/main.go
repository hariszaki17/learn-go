package main

import (
	"fmt"
)

func sayHelloWithFilter(name string, filter func(string) string)  {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string  {
	if name == "anjing" {
		return "...."
	}
	return name
}

func main()  {
	sayHelloWithFilter("mamang", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
}

/*
===NOTE===
- karena function dapat dipassing sbg paramater pada function lain

*/