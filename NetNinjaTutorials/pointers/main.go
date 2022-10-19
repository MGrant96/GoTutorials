package main

import (
	"fmt"
)

func updateName(x string) {
	x = "Michael"
}

func main() {

	var a int = 20   /* actual variable declaration */
	var ip *int      /* pointer variable declaration */
	ip = &a  /* store address of a in pointer variable*/
	fmt.Printf("Address of a variable: %x\n", &a  )
	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip )
	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip )

	name := "Mihel"
	updateName(name)
	fmt.Println("Memory address of name is: ", &name)

	// m is storing the memory location of name, the pointer
	m := &name
	fmt.Println("Memory address of m is:    ", m)
	fmt.Println("Value at memory address:   ", *m)

	fmt.Println(name)

}
