package main

import (
	"fmt"
)

func main()  {
	// const name string, akan error karna harus langsung diberi nilai
	const name string = "Komeng"
	// name = "Bobby", ini akan terjadi error
	fmt.Println(name)
	const isGood bool = true
	// tidak error ketika tidak dipanggil

	const (
		firstName = "Muhammad Diaz"
		lastName = "Zakariya"
	)
	// constant memungkinkan untuk multiple declaration
}

/*
===NOTE===
- constant adalah variable yang tidak dapat di-reassignment setelah assignment pertama
- constant wajib lansung diisi datanya setelah deklarasi contoh: const name string = "Zaki"
- constant tidak akan mengeluarkan error jika tidak dipakai
- constant tidak bisa digunakan short hand dengan syntax := tapi bisa dideklarasikan secara multiple
*/