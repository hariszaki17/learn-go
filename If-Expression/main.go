package main

import (
	"fmt"
)

func main()  {
	name := "mamank"
	class := "ini kelas"
	if name == "diaz" {
		fmt.Println("Betul namanya")
	} else if name == "joni" {
		fmt.Println("ini joni")
	} else {
		fmt.Println("gak ada yang bener gais")
	}

	if length := len(class); length < 10 {
		fmt.Println("class isinya kurang dari 10")
	} else {
		fmt.Println("class isinya lebih dari 10")
	}
}

/*
===NOTE===
- if memiliki short statement dengan sintaks => if length := len(some); length > 5
*/