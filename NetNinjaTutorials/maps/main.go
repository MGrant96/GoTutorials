package main

import (
	"fmt"
)

func main() {
	// Key Value Pairs
	// type of the key: string, type of the values: float 64
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// Looping thru the map
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// int as a key type
	phonebook := map[int]string{
		24233224: "Me",
		34432423: "CDoge",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[24233224])

	phonebook[984759373] = "bowser"
	fmt.Println(phonebook)

	phonebook[647583927] = "yoshi"
	fmt.Println(phonebook)
}
