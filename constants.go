package main

import "fmt"

const Pi = 3.14
const Greeting = "sir"

func main() {
	const World = "世界"
	fmt.Println(Greeting, World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
