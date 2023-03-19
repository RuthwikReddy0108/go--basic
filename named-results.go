package main

import "fmt"

func split(a, b, c int) (x, y, z int) {
	x = (a + b + c) * 4 / 9
	y = a - x
	z = b + c - y
	return
}

func main() {
	fmt.Println(split(10, 20, 30))
}
