package main

func checkValley(sl int, cv *int, inValley *bool) {
	if sl > 0 {
		*inValley = false
		return
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
			checkValley(sl, &cv, &inValley)
		} else if s[i] == D {
			sl--
			checkValley(sl, &cv, &inValley)
		} else {
			continue
		}
	}
	return 32
}

func main() {
	var n int32 = 8
	var s string = "UDDDUDUU"
	countingValleys(n, s)
}
