package main

import (
	"fmt"
)

func getHello(name string) string  {
	return "Hello " + name
}

func substract(num1 int, num2 int) int {
	return num1 - num2
}

func main()  {
	greet := getHello("Zaki") // memanggil fungsi getHello dan memasukan data ke dalam variable greet dgn type string
	fmt.Println(greet)
	result := substract(50, 5)
	fmt.Println("hasil pengurangan", result)
}

/*
===NOTE===
- function bisa mengembalikan data ketika dieksekusi
- untuk mendefinisikan function return value maka kita harus menulikan tipe data return valuenya
dan wajib mengembalikan data jika sudah didefinisikan
*/