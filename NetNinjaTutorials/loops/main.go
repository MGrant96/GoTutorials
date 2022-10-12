package main

import (
	"fmt"

)

func main() {
	x := 0
	for x < 5 {
		fmt.Println("Value of x is: ", x)
		x++
	}
	fmt.Println("")

	for i := 5; i > 0; i-- {
		fmt.Println("Value of i is: ", i)
	}
	fmt.Println("")

	names := []string{"mario", "luigi", "yoshi", "peach"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	fmt.Println("")
	for index, value :=range names {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}
	fmt.Println("")
	for  _, value := range names {
		fmt.Printf("The value is %v \n", value)
	}
}
