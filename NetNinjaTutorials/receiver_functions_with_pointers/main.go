package main

import (
	"fmt"
)

func main() {
	mybill := newBill("Marios Bill")
	mybill.addItem("Soup", 4.50)
	mybill.addItem("Coffee", 3.50)
	mybill.addItem("Pie", 2.50)

	mybill.updateTip(10)


	fmt.Println(mybill.format())
}
