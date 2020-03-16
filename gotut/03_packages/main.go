package main

import (
	"fmt"
	"math"

	"github.com/nickadiemus/gotut/03_packages/strutils"
)

func main() {
	// floors 2.4
	fmt.Println(math.Floor(2.4))
	// ceils 2.4
	fmt.Println(math.Ceil(2.4))

	fmt.Println(strutils.Reverse("string"))
}
