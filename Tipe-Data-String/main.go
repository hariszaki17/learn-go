package main

import (
	"fmt"
)

func main() {
	fmt.Println("Diaz") // Diaz
	fmt.Println(len("Diaz")) // 4
	fmt.Println("Muhammad Diaz") // Muhammad Diaz
	fmt.Println("Muhammad Diaz"[0]) // 77 => adalah representasi byte dari karakter M
	fmt.Println("Muhammad Diaz Zakariya") // Muhammad Diaz Zakariya
	fmt.Println("Muhammad Diaz Zakariya"[2]) // 104 => adalah representasi byte dari karakter h
}

/*
===NOTE===
- string adalah tipe data kumpulan karakter
- jumlah karakter dari nol sampai tidak terhingga
- string tidak harus mempunyai karakter bisa saja berbentuk string kosong contoh: ""
- nilai data string di go selalu dilingkupi oleh tanda petik dua ""
- function untuk string diantaranya adalah len("string") untuk mengetahui jumlah karakter string,
bisa juga untuk mengambil karakter pada sebuah string, contoh: "string"[0] => "s"
*/
