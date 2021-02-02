package main

import (
	"fmt"
)

func main()  {
	var names [3]string
	names[0] = "Muhammad" // contoh assignment pada tipe data array dengan index yang spesifik
	names[1] = "Diaz"
	names[2] = "Zakariya"
	// names[3] = "test" akan menghasilkan error array out of bound
	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	// deklarasi dan assigment array
	var values = [3]int{
		20,
		10,
		15,
	}
	fmt.Println(values)
	
	fmt.Println(len(values)) // panjang array, sesuai dengan yang dideklarasikan diawal
	values[1] = 424 // reassignment elemen array
	fmt.Println(values[1]) // menampilkan element array
}

/*
===NOTE===
- array  adalah tipe data berupa kumpulan data dengan tipe data yang sama
- saat membuat array di go kita harus mendefinisikan jumlah dari arraynya sebelum mendeklarasikannya
- pada go array tidak dapat ditambah jumlah elementnya
- untuk menambah length array dapat membuat array baru dengan length yang lebih besar
- element array pada go dimulai dari index 0
- jika mengakses index dari sebuah array diluar dari length maka akan error
- array memiliki function diantaranya adalah len(array), array[index], array[index] = value
*/