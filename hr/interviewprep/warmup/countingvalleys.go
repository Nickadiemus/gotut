package main

import "fmt"

func checkValley(sl int, cv *int, inValley *bool) {
	if sl >= 0 {
		*inValley = false
		return
	} else if sl < 0 {
		*inValley = true
	}
}

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	const U = 85
	const D = 68
	inValley := false
	sl := 0 // sea level
	cv := 0
	for i := 0; i < int(n); i++ {
		if s[i] == U {
			sl++
			prevValue := inValley
			checkValley(sl, &cv, &inValley)
			if prevValue && !inValley {
				cv++
			}
			debug(sl, cv, inValley)
		} else if s[i] == D {
			sl--
			checkValley(sl, &cv, &inValley)
			debug(sl, cv, inValley)
		} else {
			continue
		}
	}
	return int32(cv)
}

func debug(sl int, cv int, iv bool) {
	fmt.Println("________________________________________________________________")
	fmt.Println("Sea Level: \t\t\t", sl)
	fmt.Println("Counted Valleys: \t\t", cv)
	fmt.Println("inValley: \t\t\t", iv)
}

func main() {

	var s string = "UDDDUDUU"
	var n int32 = int32(len(s))
	countingValleys(n, s)
}
