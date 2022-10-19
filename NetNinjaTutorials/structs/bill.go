package main

import "fmt"

type bill struct {
	name string
	items map[string]float64
	tip float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{"Pie":2.99, "Coffee": 3.50},
		tip: 0,
	}
	return b
}

// format the bill
// bill type is recieved into this function, this is limiting it being called from a bill object only
func (b bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	// list items
	// %-20v pushes the spacing over so its all consistent
	for key, value := range b.items {
		fs += fmt.Sprintf("%-20v ... $%v \n", key+":", value)
		total += value
	}

	// total
	fs += fmt.Sprintf("%-20v ... $%0.2f", "Total:", total)
	return fs
}
