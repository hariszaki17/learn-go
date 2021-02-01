package main

import (
	"fmt"
)

func main()  {
	var name string // var => keyword, name => nama variable, string => tipe data
	name = "Muhammad Diaz Zakariya"

	// name = 1 akan error karena name dideklarasikan sbg string tapi diisi oleh integer maka error
	fmt.Println(name) // Muhammad Diaz Zakariya
	
	var frindName = "Mamat" 
	// bentuk lain dari deklarasi variable, menggunakan implicit type cast, go lang scr cerdas tau tipe data dari value
	fmt.Println(frindName)

	nameInShortHand := "Jono"
	// ini adalah bentuk shorthand declaration variable
	fmt.Println(nameInShortHand)

	var (
		firstName = "Diaz"
		lastName = "Zakariya"
	)
	// bentuk shorthand untuk deklarasi variable lebih dari 1 secara bersamaan
	fmt.Println(firstName, lastName, "2 variable")
	
	name = "Diaz Zakariya"
	// reassignment variable, data dalam variable dapat diubah"
	// tidak bisa dei reassign dengan tipe data lain spt boolean dan integer, dst
	fmt.Println(name) // Diaz Zakariya
}

/*
===NOTE===
- variable adalah tempat untuk menyimpan data
- digunakan untuk mengakses data secara berulang
- golang hanya bisa menyimpan tipe data yang sama, untuk menyimpan beberapa tipe data maka
dibutuhkan untuk membuat tipe data baru (explicit type conversion) tidak seperti js
- untuk mendeklarasikan variable dapat menggunakan keyword var lalu diikuti dengan nama variablenya
*/