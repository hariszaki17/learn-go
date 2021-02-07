package main

import (
	"fmt"
)

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main()  {
	goodbye := getGoodBye // fucction as variable
	result := getGoodBye("Jhonny") // invoke first
	fmt.Println(goodbye("Diaz"))
	fmt.Println(result)
}

/*
===NOTE===
- dalam go function adalah first class citizen, dapat dibuat sbg tipe data dan independent
- dapat disimpan di variable dan dapat dipassing sebagai value
*/