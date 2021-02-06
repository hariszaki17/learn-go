package main

import (
	"fmt"
)

func sumAllWithName(name string, numbers ...int) (nameStr string, total int)  {
	nameStr = name
	total = 0
	for _, num := range numbers {
		total += num
	}
	return
}

func sumAll(numbers ...int) (total int) {
	total = 0
	for _, number := range numbers {
		total += number
	}
	return
}

func main() {
	sum := sumAll(98, 43, 23, 543, 12, 43, 7)
	fmt.Println("ini hasil sum all", sum)

	name, total := sumAllWithName("Joni", 98, 23, 123, 43, 12)
	fmt.Printf("nama anda adalah %v dan total num adalah %v \n", name, total)
}

/*
===NOTE===
- paramaeter yang berada di posisi terakhir, memiliki  kemampuan dijadikan variable argument
- varargs artinya dapat menerima lebih dari satu input, atau semacam array
- pada variadic function paramater yang dimasukan akan di konvert menjadi array di dalam functionya
*/
