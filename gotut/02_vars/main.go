package main

import(
	"fmt"
)
/*
VARTYPES
string
bool
int
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte - uint8
rune - int32
float32 float64
complex64 complete128
*/
func main() {
	// examples
	var name string = "Nick"
	var num int = 32
	// immutable 
	const isBool = true

	// shorthand
	n := "brad"
	size := 99
	fmt.Println(name, num, isBool, n, size)
	fmt.Printf("formatted - %s\n", name)
	
}