package main

import "fmt"

func main() {
	fmt.Println("Satu =", 1)             // Satu = 1
	fmt.Println("Dua =", 2)              // Dua = 2
	fmt.Println("Tiga koma lima =", 3.5) // Tiga koma lima = 3.5
}

/*
===NOTE===
tipe data number yang ada di golang:
int8 => min: -128 max: 128
int16 => min: -32768 max: 32768
int32 => min: -2147483648 max: 2147483648
int64 => min: -9223372036854775808 max: 9223372036854775808

unsigned integer(tidak punya nilai negatif):
int8 => min: 0 max: 255
int16 => min: 0 max: 65535
int32 => min: 0 max: 4294967295
int64 => min: 0 max: 18446744073709551615

tipe data floating point(sederhananya desimal atau memiliki koma):
float32 => min: 1.18 x 10^-38 max: 3.4 x 10^38
float64 =>	min: 2.23 x 10^-308 max: 1.80 x 10^308
complex32 => complex numbers with float32 real and imaginary parts
complex64 => complex numbers with float64 real and imaginary parts

alias atau nama lain dari tipe data yang exist:
byte => uint8
rune => 32
int => minimal int32 mengikuti sistem operasi 32/64
uint => minimal uint32 mengikuti sistem operasi 32/64
*/
