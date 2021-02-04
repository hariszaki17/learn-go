package main

import (
	"fmt"
)

func main()  {
	name := "diaz"
	switch name {
	case "Eko":
		fmt.Println("ini eko")
	case "diaz":
		fmt.Println("ini diaz")
		fmt.Println("mantap gannn")
	default:
		fmt.Println("variablenya salah kali")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama benar")
	}
	
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("mantap")
	case length < 10:
		fmt.Println("kurang euyyy")
	}
}

/*
===NOTE===
- switch adalah versi sederhana dari if expression
- switch juga mendukung short statement sebelum nama varaible kondisi
*/