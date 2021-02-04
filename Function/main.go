package main

import (
	"fmt"
)

func sayHello()  {
	fmt.Println("Hello world")
}

func main()  {
	sayHello() // pemanggilan function

	for i := 0; i < 4; i++ {
		sayHello()
	}
}

/*
===NOTE==
- blok code yang berfungsi untuk digunakan berulang kali
- go hanya memiliki func tidak punya class
- setelah dibuat function dapat dipanggil untuk digunakan
*/