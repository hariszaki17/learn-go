package main

import (
	"fmt"
)

func getFullName() (string, string) { // (string, string) return value 2 data berupa string
	return "Diaz", "Zakariya"
}

func main()  {
	firstName, lastName := getFullName() // returning value ditempatkan di 2 buah variable
	firstNameSaja, _ := getFullName() // returning value ditempatkan di 2 buah variable
	fmt.Println("ini adalah firstName", firstName)
	fmt.Println("ini adalah lastName", lastName)
	fmt.Println("ini adalah firstName cuma ngambil 1 value saja", firstNameSaja)
}

/*
===NOTE===
- pada golang function dapat mengembalikan data lebih dari 1
- untuk memerintah function untuk mengembalikan multiple value maka harus
didefinisikan dulu tipe data
*/