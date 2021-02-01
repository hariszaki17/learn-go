package main

import (
	"fmt"
)

func main()  {
	var a int = 10
	var b int = 10
	var c int = a + b
	fmt.Printf("hasil dari penjumlahan a:%v dan b:%v adalah %v \n",a, b, c)

	a += 40
	b -= 5
	c *= 2
	// ketiga operator ini adalah augmented assignment atau pengoperasian sekaligus reassignment
	fmt.Printf("hasil augmented a:%v, b:%v, c:%v \n", a, b, c)
	
	var unaryA int = 5
	var unaryB int = 3
	var unaryBool bool = false

	unaryA++
	unaryB--
	unaryBool = !unaryBool
	fmt.Println("ini unaryA increment", unaryA)
	fmt.Println("ini unaryB decrement", unaryB)
	fmt.Println("ini unaryBool converse", unaryBool)

	unaryA = -unaryA
	unaryB = -20
	fmt.Println("ini unaryA negatif", unaryA)
	unaryB = -unaryB // minus dan minus = plus
	fmt.Println("ini unaryB positif", unaryB)
} 