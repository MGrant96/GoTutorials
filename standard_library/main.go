package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	fmt.Println("Strings Library")
	greeting := "This is a Greeting!"
	fmt.Println(strings.Contains(greeting, "is"))
	fmt.Println(strings.ReplaceAll(greeting, "Greeting", "Great Greeting"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "a"))
	fmt.Println(strings.Split(greeting, " "))
	// Original value is unchanged
	fmt.Println("Original string value = ", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	fmt.Println("Ages Array       ", ages)
	sort.Ints(ages)
	fmt.Println("Sorted Ages Array", ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	fmt.Println("Names", names)
	fmt.Println(sort.SearchStrings(names, "bowser"))

}
