package main

import "fmt"

func main() {

	var nameOne string = "Mario"
	var nameTwo = "Luigi"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameThree = "Peach"
	nameFour := "Yoshi"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	var numOne int8 = 127
	var numTwo int8 = -128
	var numThree uint8 = 24

	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 534543543345345345.345345
	scoreThree := 1.5

	fmt.Println(scoreThree + scoreTwo - float64(scoreOne))
}
