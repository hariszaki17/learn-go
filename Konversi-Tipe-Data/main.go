package main

import (
	"fmt"
)

func main()  {
	// batas atas value yang dapat ditampung int8 adalah 127
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32) // 100000 output benar
	fmt.Println(nilai64) // 100000 output benar
	fmt.Println(nilai8) // -96 output salah

	var nilai32True int32 = 127
	var nilai64True int64 = int64(nilai32True)
	var nilai8True int8 = int8(nilai32True)

	fmt.Println(nilai32True) // 127 output benar
	fmt.Println(nilai64True) // 127 output benar
	fmt.Println(nilai8True) // 127 output benar

	var nilai32False int32 = 128
	var nilai64False int64 = int64(nilai32False)
	var nilai8TrueFalse int8 = int8(nilai32False)

	fmt.Println(nilai32False) // 128 output benar
	fmt.Println(nilai64False) // 128 output benar
	fmt.Println(nilai8TrueFalse) // -128 output salah, kembali ke titik paling rendah

	// konversi byte ke string
	name := "diaz"
	dByte := name[0]
	fmt.Println(dByte) // 100
	dString := string(dByte)
	fmt.Println(dString) // "d" setelah dikonversi
}

/*
===NOTE===
- karna golang tidak dapat melakukan implicit type cast maka diperlukan metode untuk
mengubah sebuah variable dari satu tipe data ke tipe data lain
- jika dilihat pada konversi dari int32 ke int 8 terjadi integer overflow
yang mana int8 hanya memiliki range value antara -128 sampai 128 sedangkan value yang dikonversi
bernilai 100000 hal ini akan terjadi pengulangan dari titik paling bawah

*/