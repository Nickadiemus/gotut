package main

import "fmt"

func helloText(s string) string {
	return ("Hello " + s)
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(helloText("Nick"))
}
