package main

import (
	"fmt"
)

func main()  {
	var a int = 20
	var b int = 10
	var isGreater bool
	isGreater = a > b
	fmt.Println("a greater than b is", isGreater)

	var name1 string = "diaz"
	var name2 string = "diaz"
	var isSame bool = name1 == name2
	fmt.Println("name1 equals name2 is", isSame)
	name2 = "DIAZ"
	isSame = name1 == name2
	fmt.Println("name1 equals name2 with different case is", isSame)

	var compLetterA string = "ab"
	var compLetterB string = "ac"
	fmt.Println(compLetterA < compLetterB) 
	// komparasi seperti ini bisa digunakan untuk melakukan sorting pada tipe data string

	var value1 int = 10
	var value2 int = 20
	fmt.Println(value1 > value2) // false
	fmt.Println(value1 < value2) // true
	fmt.Println(value1 == value2) // false
	fmt.Println(value1 != value2) // true
}

/*
===NOTE===
- operasi perbandingan adalah operasi untuk membandingkan dua buah data
- operasi ini akan menghasilkan nilai boolean true or false
- jika hasil operasi benar maka nilainya adalah true
- jika hasil operasi salah maka nilai adalah false
- operasi perbandingan diantaranya >, <, >=, <=, ==, !=
*/