package main

import (
	"fmt"
)

func main()  {
	counter := 1
	for counter <= 10 {
		fmt.Println("ini counter:", counter)
		counter++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("berhasil gais:", i)
	}

	arr := [10]int{1, 2, 4, 5, 6, 7, 8, 9, 1}
	for index, val := range arr {
		fmt.Println("ini elemetnya", val)
		fmt.Println("ini indexnya", index)
	}

	mapData := map[string]string{
		"nama": "diaz",
		"position": "developer",
	}

	for key, val := range mapData {
		fmt.Printf("ini key: %v dan ini data: %v \n", key, val)
	}
}

/*
===NOTE===
- dalam go loops hanya memiliki 1 bentuk yaitu for
- for memiliki 2 statement yaitu init statement dan post statement
contoh int := 0 init statement, int++ post statement
- perulangan dapat digunakan juga untuk iterasi map, slice, atau array
*/