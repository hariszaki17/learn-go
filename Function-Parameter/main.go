package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
)

// func sayHello(firstName string, lastName string)  {
// 	fmt.Println("hello", firstName, lastName)
// }

func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("simple shell")
	fmt.Println("------------------")
	fmt.Println("-----Pilih menu makanan-------")
	fmt.Println("1. Baso")
	fmt.Println("2. Ayam")
	fmt.Println("3. Nasi")
	fmt.Println("4. Roti")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	if text == "1" {
		fmt.Println("harga baso 10000 dengan pilihan:", text)
	} else if text == "2" {
		fmt.Println("harga ayam 20000 dengan pilihan:", text)
	} else if text == "3" {
		fmt.Println("harga nasi 5000 dengan pilihan:", text)
	} else if text == "4" {
		fmt.Println("harga roti 4000 dengan pilihan:", text)
	}
	// sayHello("haris", "")
}

/*
===NOTE===
*/