package main

import "fmt"

func checkPos(position int32, c []int32, jumps *int) int32 {
	// //check out of bounds
	// fmt.Println("Comparing: ")
	// fmt.Printf("%d > %d || %d >= %d\n", (position + 2), int32(len(c)), (position + 1), int32(len(c)))

	oneJump := true
	twoJump := true

	if (position + 2) >= int32(len(c)) {
		twoJump = false
	}
	if (position + 1) >= int32(len(c)) {
		oneJump = false
	}

	if (twoJump) && c[position+2] == 0 {
		*jumps++
		return position + 2
	} else if (oneJump) && c[position+1] == 0 {
		*jumps++
		return position + 1
	} else {
		return -1
	}

}

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {

	jumps := 0
	var curpos int32 = 0
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
	return int32(jumps)
}

func main() {
	//tests
	var game1 = []int32{0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0}
	var game2 = []int32{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}
	var game3 = []int32{0, 0, 0, 1, 0, 0}

	fmt.Println(jumpingOnClouds(game1))
	fmt.Println(jumpingOnClouds(game2))
	fmt.Println(jumpingOnClouds(game3))
}
