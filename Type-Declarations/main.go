package main

import (
	"fmt"
)

func main()  {
	type name string
	type finished bool
	// syntax di atas berfungsi untuk membuat tipe data baru representasi dari tipe data string
	var firstName name = "Bob"
	// penggunaan tipe data baru hasil dari type declaration
	fmt.Println(firstName) // Bob
	num := "11111111"
	fmt.Println(name(num)) // tipe data num berubah dari integer menjadi name(string)
	var isFinished finished = false
	fmt.Println(isFinished)
}

/*
===NOTE===
- kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
- berfungsi sebagai alias untuk memudahkan coding
*/