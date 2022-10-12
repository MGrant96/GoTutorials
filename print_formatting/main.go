package main

import "fmt"

func main() {
	age := 26
	name := "Michael"
	// Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("Newline")
	// Println
	fmt.Println("Hello There!")
	fmt.Println("General Kenobi!")
	fmt.Println("My age is ", age, " and my name is ", name, "!")
	// Formatted Printf
	fmt.Printf("My age is %v and my name is %v!\n", age, name)
	fmt.Printf("Age is of type %T!\n", age)
	fmt.Printf("You scored %0.2f points!\n", 255.55555)
	// Sprintf Save formatted Strings
	var str = fmt.Sprintf("My age is %v and my name is %v!\n", age, name)
	fmt.Println("The Saved string is: ", str)
}
