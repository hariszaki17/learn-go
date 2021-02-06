package main

import (
	"fmt"
)

func getCompleteName() (firstName, lastName string)  {
	firstName = "Mamang"
	// lastName = "Jojon"
	return
}

func main()  {
	firstName, lastName := getCompleteName()
	fmt.Println("ini first name", firstName)
	fmt.Println("ini last name", lastName)
}

/*
===NOTE===
- kita dapat membuat variable secara langsung di tipe data return sebuah function
*/