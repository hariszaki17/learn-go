package main

import (
	"fmt"
)

func main()  {
	person := map[string]string{
		"name": "joko",
		"address": "depok",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	person["title"] = "programmer"
	fmt.Println(person)
	animal := make(map[string]string)
	animal["type"] = "mamal"
	fmt.Println(animal)
	delete(person, "title")
	fmt.Println(person)
	fmt.Println(len(person))
}

/*
===NOTE===
jenis function pada map
- len(map) => mendapatkan jumlah data di map
- map[key] => mendapatkan value berdasarkan key
- map[key] = value => assignment untuk key pada map
- make(map[typekey]typeValue) => membuat map
- delete(map, key) => menghapus data pada map berdasarkan key
*/