package main

import "fmt"

func main() {

	color := "red"

	// if cond
	if color == "red" {
		fmt.Println("The color is " + color)
	} else {
		fmt.Println("The color is not red!")
	}

	color = "blue"
	if color == "blue" {
		fmt.Println("The color is " + color)
	} else if color == "red" {
		fmt.Println("The color is " + color)
	} else {
		fmt.Println("The color was neither blue or red!")
	}
	color = "black"
	switch color {
	case "red":
		fmt.Println("The color is red!")
	case "blue":
		fmt.Println("The color is blue!")
	default:
		fmt.Println("The color was neitehr blue or red!")
	}
}
