package main

import "fmt"

func main() {
	arr := []string{"Hello", "Good", "Sir"}

	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	var arrs [3]string

	arrs[0] = "King"
	arrs[1] = "Alfred"
	arrs[2] = "III"

	fmt.Println(arrs[0:3])
}
