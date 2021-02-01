package main

import (
	"fmt"
)

func main()  {
	var avg int = 80
	var cognitiveMarks = 95
	var presenceMarks = 80

	var isPassedCognitive bool = cognitiveMarks >= avg
	var isPassedPresence bool = presenceMarks >= avg
	var isPassedFinal = isPassedCognitive && isPassedPresence
	fmt.Println("hasil akhir adalah", isPassedFinal)
}

/*
===NOTE===
tabel operasi &&
-----------------
nilai1 | operator | nilai2 | hasil
true   |    &&    | true   | true
true   |    &&    | false  | false
false  |    &&    | true   | false
false  |    &&    | false  | false


tabel operasi ||
-----------------
nilai1 | operator | nilai2 | hasil
true   |    ||    | true   | true
true   |    ||    | false  | true
false  |    ||    | true   | true
false  |    ||    | false  | false

*/