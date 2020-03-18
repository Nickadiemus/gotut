package main

import (
	"fmt"
	"math"
)

func unique(li []int32) []int32 {
	//map
	deduplist := make(map[int32]bool)
	list := []int32{}
	// {10, 20, 20, 10, 10, 30, 50, 10, 20}
	for _, key := range li {
		fmt.Println(key)
		if _, value := deduplist[key]; !value { // Key: 10 Value: false
			deduplist[key] = true
			list = append(list, key)
		}

	}
	return list
}

// Complete the sockMerchant function below.
// n: number of socks in the pile
// ar: colors of each sock
func sockMerchant(n int32, ar []int32) int32 {
	if n < 1 && 100 < n {
		return 0
	}
	var matches int32 = 0
	pairs := make(map[int32]int)
	uniquelist := unique(ar)

	for _, value := range uniquelist {
		pairs[value] = 0
	}
	for _, value := range ar {
		if _, p := pairs[value]; p {
			pairs[value]++
		}
	}

	for _, v := range pairs {
		matches += int32((math.Floor(float64(v) / 2)))
	}

	return matches
}

func main() {
	var n int32 = 9
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	sockMerchant(n, ar)

}
