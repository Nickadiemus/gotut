package main

import "fmt"

func checkPos(position int32, c []int32, jumps *int) int32 {
	//check out of bounds
	fmt.Println("Comparing: ")
	fmt.Printf("%d > %d || %d >= %d\n", (position + 2), int32(len(c)), (position + 1), int32(len(c)))

	// if (position + 2) > int32(len(c)-1) {
	// 	fmt.Println("Matched Pos")
	// 	return -1
	// } else if (position + 1) > int32(len(c)-1) {
	// 	fmt.Println("Matched Pos")
	// 	return -1
	// } else if position == (int32(len(c) - 1)) {
	// 	if c[position-1] == 0 {
	// 		fmt.Println("Matched End Case")
	// 		fmt.Printf("%d == 0\n", c[position-1])
	// 		*jumps++
	// 	}
	// 	return -1
	// } else if c[position+2] == 0 {
	// 	fmt.Println("Matched Pos + 2")
	// 	fmt.Printf("%d == 0\n", c[position+2])
	// 	*jumps++
	// 	return position + 2
	// } else if c[position+1] == 0 {
	// 	fmt.Println("Matched Pos + 1")
	// 	fmt.Printf("%d == 0\n", c[position+1])
	// 	*jumps++
	// 	return position + 1
	// } else {
	// 	fmt.Println("Matched Pos + 0")
	// 	return position
	// }

	if (position+2 >= int32(len(c))) || (position+1 >= int32(len(c))) || (position >= int32(len(c))) {
		if position == int32(len(c)) {
			*jumps++

		}
		return -1
	}

	// oneJump := false
	// twoJump := false

	return -1
	// if c[position+2]

}

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {

	jumps := 0
	var curpos int32 = 9
	var iterator int32 = 0
	for iterator < int32(len(c)) {
		iterator++
		curpos = checkPos(curpos, c, &jumps)
		fmt.Println("CurrentPos: ", curpos)
		fmt.Println("CurJumps: ", jumps)
		fmt.Println("")
		if curpos != -1 {
			continue
		} else {
			return int32(jumps)
		}
	}
	return int32(-1)

}

func main() {
	// var game1 = []int32{0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0}
	var game2 = []int32{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}
	// var game3 = []int32{0, 0, 0, 1, 0, 0}

	// fmt.Println(jumpingOnClouds(game1))
	fmt.Println(jumpingOnClouds(game2))
	// fmt.Println(jumpingOnClouds(game3))
}
