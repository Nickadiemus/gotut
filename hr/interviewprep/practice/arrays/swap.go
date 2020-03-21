package main

import (
	"fmt"
	"sort"
)

// type swap

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {

	swaps := 0 // number of minimum swaps
	/*
		since golang map doens't support
		indexed maps I had to create a map
		that indexed the maps keys after the
		sortkey was created. This aided in
		determining the minimum swaps needed
	*/
	revarrpos := make(map[int]int)
	/*
		mapped array
	*/
	arrpos := make(map[int]int) // mapped arr indexes
	for i := 0; i < len(arr); i++ {
		arrpos[int(arr[i])] = i
	}
	sortkey := make([]int, 0)
	for key := range arrpos {
		sortkey = append(sortkey, key)
	}
	hasVis := make([]bool, len(arrpos))

	// sort the key values
	sort.Ints(sortkey)
	for index, k := range sortkey {
		revarrpos[index] = k
	}
	for index, k := range sortkey {
		if hasVis[index] || arrpos[k] == index {
			continue
		}

		cycle := 0
		j := index
		for !hasVis[j] {
			hasVis[j] = true
			j = arrpos[revarrpos[j]]
			cycle++
		}
		if cycle > 0 {
			swaps += (cycle - 1)
		}
	}
	// fmt.Println(arrpos)
	return int32(swaps)
}

func main() {
	// fmt.Println("swap.go program")
	// arr := []int32{1, 3, 5, 2, 4, 6, 7}
	arr := []int32{4, 3, 1, 2}
	fmt.Println(minimumSwaps(arr))
}
