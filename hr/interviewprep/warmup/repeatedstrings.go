package main

import (
	"fmt"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	var c int64 = 0
	var lc int64 = 0
	var calc int64
	var achar uint8 = 97
	sfactor := (n % int64(len(s)))
	ssfactor := (n / int64(len(s)))
	if len(s) == 1 && s[0] == achar {
		return n
	}
	for i := 0; i < len(s); i++ {
		if s[i%len(s)] == achar {
			c++
		}
	}
	for i := 0; i < int(sfactor); i++ {
		if s[i%len(s)] == achar {
			lc++
		}
	}
	calc = c*ssfactor + lc

	return calc
}

func main() {
	// test driver
	s1 := "kmretasscityylpdhuwjirnqimlkcgxubxmsxpypgzxtenweirknjtasxtvxemtwxuarabssvqdnktqadhyktagjxoanknhgilnm"
	var n int64 = 736778906400

	fmt.Println(repeatedString(s1, n))
}
