package main

import (
	"fmt"
)

func main() {
	mybill := newBill("Marios Bill")

	fmt.Println(mybill)

	fmt.Println(mybill.format())
}
