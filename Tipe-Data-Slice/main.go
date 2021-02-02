package main

import (
	"fmt"
)

func main()  {
	var values = [10]int{ 1, 3, 5, 7, 3, 1, 4, 1, 2, 8 }
	var slicesInt = values[2:6]
	fmt.Println(slicesInt)
	fmt.Println(cap(slicesInt))

	var months = [12]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	var monthSlice = months[4:8]
	fmt.Println(cap(monthSlice)) // output 8, selisih len(months) dengan pointer pertama(4)
	fmt.Println(len(monthSlice)) // output 4, selisih pointer pertama dengan pointer kedua
	months[4] = "MEI"
	monthSlice = append(monthSlice, "tambahan")
	fmt.Println(monthSlice)
	// output [MEI juni juli agustus], ketika mengubah data arary maka slice akah ikut berubah slice 1 pointer memory dengan array parent
	// slice pada dasarnya refer ke array, semua slice yang refer ke array akan berubah ketika array berubah
	fmt.Println(months)
	// jika melakukan append pada slice maka data array juga akan berubah
	
	implicitLength := [...]string{ "joni", "supardi", "mamang", "padul", }
	// jumlah slice mengikuti jumlah element
	fmt.Println(implicitLength)
	fmt.Println(len(implicitLength))

	var usingMake = make([]int, 4, 6)
	fmt.Println(usingMake)
	fmt.Println(cap(usingMake))

	var contohArray = [5]int{ 1, 5, 2, 5, 4 }
	var contohSlice = contohArray[3:]
	contohSlice = append(contohSlice, 9)
	contohSlice = append(contohSlice, []int{ 3, 5, 6 }...)
	/* 
	append jika dilakukan pada slice dengan refer pointer sejumlah dengan 
	cap array maka slice akan membuat data array baru dan tidak mengubah data awal array
	*/
	fmt.Println(contohSlice)
	fmt.Println(len(contohSlice))
	fmt.Println(cap(contohSlice))
	fmt.Println(contohArray)

	var cobaMake = make([]int, 4, 5)
	fmt.Println(cobaMake)
	copy(cobaMake, []int{3, 5, 2, 9})
	fmt.Println(cobaMake)
	cobaMake = append(cobaMake, 8)
	cobaMake = append(cobaMake, 9)
	cobaMake = append(cobaMake, 10)
	fmt.Println(cobaMake)
}

/*
===NOTE===
- tipe data slice adalah potongan dari tipe data array
- perbedaan slice dengan array adalah slice bisa berubah ukuran
- slice dan array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian
atau seluruh data di array
- data slice memiliki 3 komponen, yaitu pointer, length dan capacity sementara array hanya memiliki length
- untuk membuat slice pertama harus membuat array terlebih dahulu
dan pada dasarnya slice ini menggunakan data yang berada di array
- serta slice punya capacity yang mana cap ini tidak boleh melebih dari length
- slice kapasitasnya bisa kita perbesar hanya sampai nilai dari fungsi cap()
- cap pada slice didapatkan dari pointer awal sampai dengan length dari array yang dipakai,
contoh array dengan length 10, slice = array[4:8] maka cap dari slicenya adalah 6 karena selisih
10 dan pointer pertama(4) adalah 6

===PEMBUATAN SLICE===
array[low:high] => membuat slice dari array dimulai index low sampai index sebelum high
arrat[low:] => membuat slice dari array dimulai index low sampai index akhir di array
array[:high] => membuat slice dari array dimulai index 0 hingga index sebelum high
array[:] => membuat slice dari array dimulai index 0 hingga index akhir

===WARNING===
- perhatikan deklaraasi array, untuk implisit length array bisa dilakukakndengan cara:
value := [...]int{1, 3}, sementara slice dengan cara value := []int{1, 3}
- 
*/