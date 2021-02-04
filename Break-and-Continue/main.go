package main

import (
	"fmt"
)

func main()  {
	for i := 0; i < 10; i++ {
		if i <= 5 {
			fmt.Println("ini belum break", i)
		} else {
			break
		}
	}

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("perulangan ke", i)
	}
}

/*
===NOTE===
*/