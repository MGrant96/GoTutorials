package main

import "fmt"

func main() {
	var ages [3]int = [3]int{20, 25, 30}
	var newAges = [3]int{22, 27, 32}
	fmt.Println(ages, len(ages))
	fmt.Println(newAges, len(newAges))

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	fmt.Println(names, len(names))
	names[1] = "luigi"
	fmt.Println(names, len(names))

	// Slices
	var scores = []int{100, 50, 60}
	fmt.Println(scores)
	scores[2] = 25
	fmt.Println(scores)
	scores = append(scores, 85)
	fmt.Println(scores)

	// Slice ranges
	// inclusive of the first number but not the
	// second, so gets the 1 and 2 element of names
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, "\n", rangeTwo, "\n", rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
}
