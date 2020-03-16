package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("b: ", *b)

	// change val with pointer
	*b = 10
	fmt.Println("*b: ", a)
}
