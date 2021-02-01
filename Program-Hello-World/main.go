package main
/* package main adalah tempat dimana dilakukan initial invocation 
yang paling tidak harus dimiliki pada suatu program */

import "fmt"
/* import untuk menggunakan module library atau local module yg kita buat
sedangkan fmt adalah library bawaan dari go yang fungsinya untuk input dan output stream
hampir sama dengan console.log pada js */

func main()  {
	fmt.Println("Hello World")
}
/* pada package main harus terdapat fungsi main sbg fungsi 
yang akan diinvoke pertama kali */

/* untuk menjalankan go file dapat digunakan commang compile yaitu: go build main.go
yang mana hal ini akan melakukan kompilasi dengan output file native binary sesuai dengan
nama file yang dikompilasi contoh file => main sedangkan untuk development kita dapat
menggunakan command go run main.go dan ini sangat mudah karna tidak perlu mejalankan ulang
file hasil kompilasi karena go run sudah otomatis menjalankan file yang dipilih */

/* untuk menjalankan file hasil kopmilasi dengan cara ./main pada linux */
