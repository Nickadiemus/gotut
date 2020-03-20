package main

import "fmt"

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	len := int32(len(a))
	var newArray []int32
	for i := int32(0); i < d; i++ {
		newArray = a[1:len]
		newArray = append(newArray, a[0])
		a = newArray
	}
	return a
}

func main() {
	arr := []int32{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51}
	d := int32(10)

	fmt.Println(rotLeft(arr, d))
}
