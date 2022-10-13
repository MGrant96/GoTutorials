package main

import (
	"fmt"
)

func updateName(x string) {
	x = "Wedge"
}

func updateName2(x string) string {
	x = "Michael"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 3.5
}
func main() {
	// Non-Pointer Values
	// Group A Types: Strings, Ints, Bools, Floats, Arrays, Structs
	name := "tifa"
	updateName(name) // Only the copy is changed
	fmt.Println(name) // Will still return "tifa"

	// name is now going to be equal to what is returned by the new function
	name = updateName2(name)
	fmt.Println(name)

	// Pointer Wrapper Values
	// Group B Types: Slices, Maps, Functions
	menu := map[string]float64 { 
		"pie": 4,
		"ice cream": 2.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
