package main

import "fmt"

func main() {
	r := []int{23, 25, 65, 4, 35, 776, 454, 10, 78}

	for i, rnum := range r {
		fmt.Printf("%d - rnum %d\n", i, rnum)
	}

	// not using index
	for _, rnum := range r {
		fmt.Printf("rnum %d\n", rnum)
	}

	// add rnums together
	sum := 0

	for _, rnum := range r {
		sum += rnum
	}

	fmt.Println("Total:", sum)

	// range with map
	emails := map[string]string{"Bob": "Bob@gmail.com", "Shannon": "shannon@gmail.com", "Mike": "mike@gmail.com"}

	for k, v := range emails {
		fmt.Println(k, v)
	}
}
