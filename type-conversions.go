package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 10
	var f = math.Sqrt(x*x + y*y)
	var z = f
	fmt.Println(x, y, z)
}
